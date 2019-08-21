countcodes() {
  echo $( curl -s localhost:8288/codes/ | jq '. | length' 2>&1 )
}

# Make sure we have our testing directory to mount in docker
mkdir -p testing

# Empty JSON to start with
echo '{}' > testing/ports.json

# Start our services with our testing JSON volume
docker-compose -f docker-compose.yml -f testing.yml up -d $FORCEBUILD

# Count how many items we have
c=$(countcodes)
if [ $c -ne 0 ]; then
  echo "Empty JSON file has non-zero shortcodes"
fi

cat >testing/ports.json <<ONEPORT
{ "AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [], "regions": [], "coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai", "unlocs": [ "AEAJM" ], "code": "52000" } }
ONEPORT

curl -s localhost:8288/reload/ >/dev/null
if [ $? -ne 0 ]; then
  echo "Reloading failed"
fi

c=$(countcodes)
if [ $c -ne 1 ]; then
  echo "One port JSON file doesn't have 1 shortcode"
fi

curl -s localhost:8288/reload/ >/dev/null
if [ $? -ne 0 ]; then
  echo "Reloading failed"
fi

c=$(countcodes)
if [ $c -ne 1 ]; then
  echo "One port JSON file reloaded doesn't have 1 shortcode"
fi

cat > testing/ports.json <<DUPLICATED
{ "AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [], "regions": [], "coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai", "unlocs": [ "AEAJM" ], "code": "52000" },
"AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [], "regions": [], "coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai", "unlocs": [ "AEAJM" ], "code": "52000" } }
DUPLICATED

curl -s localhost:8288/reload/ >/dev/null
if [ $? -ne 0 ]; then
  echo "Reloading failed"
fi

c=$(countcodes)
if [ $c -ne 1 ]; then
  echo "Duplicated port JSON file reloaded doesn't have 1 shortcode"
fi


cat > testing/ports.json <<TWOPORTS
{ "AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [],
    "regions": [], "coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai",
    "unlocs": [ "AEAJM" ], "code": "52000" }, "AEAUH": { "name": "Abu Dhabi", "coordinates": [ 54.37,
      24.47 ], "city": "Abu Dhabi", "province": "Abu Z¸aby [Abu Dhabi]", "country": "United Arab Emirates",
      "alias": [], "regions": [], "timezone": "Asia/Dubai", "unlocs": [ "AEAUH" ], "code": "52001" } }
TWOPORTS

curl -s localhost:8288/reload/ >/dev/null
if [ $? -ne 0 ]; then
  echo "Reloading failed"
fi

c=$(countcodes)
if [ $c -ne 2 ]; then
  echo "Two ports JSON file reloaded doesn't have 2 shortcodes"
fi

# Ideally we'd be able to test going back to one port but deletion isn't implemented
json=$(curl -s localhost:8288/shortcode/AEAUH)
data=$(echo "$json" | jq -r '[.name, .country, .timezone, .code, .shortcode] | join("|")')
if [ "$data" != "Abu Dhabi|United Arab Emirates|Asia/Dubai|52001|AEAUH" ]; then
  echo "Fetching by shortcode failed: $data"
fi

# Stop our services
docker-compose -f docker-compose.yml -f testing.yml down

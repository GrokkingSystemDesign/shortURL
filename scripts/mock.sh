curl -L -s $(
    curl -X POST \
    http://localhost:8080/ \
    -H 'content-type: application/json' \
    -d '{"long": "https://github.com/GrokkingSystemDesign/shortURL"}'
)
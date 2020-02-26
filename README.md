# Smithsonian-image-scraper
a small scraper to get images from the smithsonian media collection based on search terms
Please don't use this to download the entire site. If you want to do that, just get the full catalog from the [smithsonian github](https://github.com/Smithsonian/OpenAccess) and parse those files instead of slamming their api server.

You'll need a data.gov api key to use this. Please get one at [The Data.gov site](https://api.data.gov/signup/)


usage: go run scrape.go --apikey $API_KEY term1 term2 ... termn

#!/bin/bash

#cd /Users/philip/go/src/github.com/truckcoinswap/bol-app/app-bol/public/img/icons
cd /Users/philip/go/src/github.com/pschlump/gatus/web/app/src/assets

# Step 1 - convert favicon.svg -> 1080x1080 png.

rsvg-convert -w 1080 logo.svg >logo.png
cp logo.png /tmp

# convert favicon-1080x1080.png -resize '192x192' android-chrome-192x192.png
# convert favicon-1080x1080.png -resize '512x512' android-chrome-512x512.png
# convert favicon-1080x1080.png -resize '154x154' -gravity center -background transparent -extent 192x192 android-chrome-maskable-192x192.png
# convert favicon-1080x1080.png -resize '410x410' -gravity center -background transparent -extent 512x512 android-chrome-maskable-512x512.png
# convert favicon-1080x1080.png -resize '180x180' -background white apple-touch-icon.png
# convert favicon-1080x1080.png -resize '60x60' -background white apple-touch-icon-60x60.png
# convert favicon-1080x1080.png -resize '76x76' -background white apple-touch-icon-76x76.png
# convert favicon-1080x1080.png -resize '120x120' -background white apple-touch-icon-120x120.png
# convert favicon-1080x1080.png -resize '180x180' -background white apple-touch-icon-180x180.png
# convert favicon-1080x1080.png -resize '152x152' -background white apple-touch-icon-152x152.png
# convert favicon-1080x1080.png -resize '16x16' favicon-16x16.png
# convert favicon-1080x1080.png -resize '32x32' favicon-32x32.png
# convert favicon-1080x1080.png -resize '144x144' msapplication-icon-144x144.png
# convert favicon-1080x1080.png -resize '150x150' -background transparent -compose Copy -gravity center -extent 270x270 mstile-150x150.png

cd /Users/philip/go/src/github.com/pschlump/gatus/web/app/public

convert /tmp/logo.png -resize '180x180' apple-touch-icon.png
convert /tmp/logo.png -resize '16x16' favicon-16x16.png
convert /tmp/logo.png -resize '32x32' favicon-32x32.png
convert /tmp/logo.png -resize '192x192' logo-192x192.png
convert /tmp/logo.png -resize '512x512' logo-512x512.png

convert favicon-32x32.png favicon.ico



# Discord Last.FM Scrobbler

This script will update your Discord "Listening to" status with whatever you are listening to according to Last.FM

![Discord Screenshot](https://i.imgur.com/FflYUk0.png)

Setup is fairly straightforward and only takes a few minutes. You'll need two things:

## **1. Last.FM API Key**

Head over to the [Last.FM API Page](https://www.last.fm/api/account/create) and sign in with your existing Last.FM username and password. It should bring you to the **Create API account** page and ask you for a few things.

It doesn't really matter what you put in most of the fields, but it should probably look something like this:

![LastFM Create API Account Screenshot](https://i.imgur.com/VQYa8nr.png?1)

After clicking Submit you should get a confirmation page with two items: *API Key* and *Shared Secret*. The API Key is the only one you need for this, but I recommend you save both for future use just in case, as they don't actually provide a way to retrieve these later.

![LastFM API Account Created Screenshot](https://i.imgur.com/oQTdNgX.png)

Copy and paste the API Key value into the config file in the `api_key = xxx` line

## **2. Discord User Token**

For this one you'll need to use the Desktop or Web app - it will not work on mobile.

If you are using the desktop app:

- Press **Ctrl+Shift+I**
- Click the "*Application*" tab
- Click and expand the "*Local Storage*" section
- Click on the only entry in this section, "*https://discordapp.com*"
- Right click -> Edit Value in the field to the right of "*token*"
- Copy and paste the token value into the config file on the `token = xxx` line and remove the quotation marks from it.

![Desktop Token](https://i.imgur.com/EEN2mnv.png)

If you are using Discord in a browser:

- Press **F12**
- Click the "*Storage*" tab
- Click and expand the "*Local Storage*" section
- Click on the only entry, "*https://discordapp.com*"
- Copy the value beside the "*token*" entry and paste it into your config file without the quotation marks.

![Browser Token](https://i.imgur.com/OFrhTHE.png)

## When you're done

Save your config file as "*config.ini*" and it should look something like this:

![Finished Config File](https://i.imgur.com/X1pO6Di.png)

Now just run the executable. It should connect to Discord and immediately start setting your "*Playing*" status to whatever you're listening to on Last.FM

If it's working, it will look like this:

![Running Executable](https://i.imgur.com/lsb0GFx.png)

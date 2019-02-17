# Pick
Built with [Golang](https://golang.org) and [Google Cloud Functions](https://cloud.google.com/blog/products/application-development/cloud-functions-go-1-11-is-now-a-supported-language).

**Pick** is the best bot to help you making a quick decision from different choices in Slack.

Hard to make a quick decision? Don't want to waste your time to decide something you don't really care about?
* Where to go for team lunch?
* What should we organzie for next team-building activity?
* Who should be the one to lead this project?
* ..... more questions that you don't want to put too much effort on it

Simply ask `/pick` to help you picking one from the choices. Even you can use it for **lucky draw**, to see who are in lucky too!

More features are coming soon, stay tuned!

For any feedback or support, please email to [pick@andrewmmc.com](pick@andrewmmc.com)

## Usage
### /pick [choice] [choice]...
Example:

`/pick Chicken Pizza Kebab Pasta Rice`

`/pick Andrew Ignatius Jake Jay Divya`

**Pick** will answer you the choice from the list you provided.

### /pick help
Show help infobox for the commands usage.

## Test
WIP

## Deploy
For reference only, basically just install the application on Slack and ignore the below part :)

If you have no latest `gcloud` beta commands installed, run the following:
```
gcloud components update
gcloud components install beta
```

Given that you installed the latest `gcloud` beta commands,
```
gcloud beta functions deploy install --env-vars-file .env.yaml --runtime go111 --entry-point Install --trigger-http --region asia-northeast1
gcloud beta functions deploy authCallback --env-vars-file .env.yaml --runtime go111 --entry-point AuthCallback --trigger-http --region asia-northeast1
gcloud beta functions deploy getAnswer --env-vars-file .env.yaml --runtime go111 --entry-point GetAnswer --trigger-http --region asia-northeast1
```

## Author
* [Andrew Mok](https://andrewmmc.com) (@andrewmmc)

## Support
* Give this repo a **star** if you like :)
* For any questions, please feel free to [open an issue here](../../issues) or email to [pick@andrewmmc.com](pick@andrewmmc.com).

## Donations
* We are providing **free service** to support any basic usage from the users, meaning that no charge if you installed this application. 
* If you would like to support the continuous maintenance of this project, please [**feel free to donate on PayPal**](https://www.paypal.me/andrewmmc). Your donation is highly appreciated, thank you!
## qTips

Right now the only example here is a basic **window** with a **widget** container with a **form** and a **button** inside it.

The form consists of rows of corresponding labels and inputs. The button does nothing as of writing this readme.

## Prerequisites
**QT**

Version 5.12.0 was used to compile this project, it was the only one that worked with the goqt bindings in my experience. You can download this version from [here](https://download.qt.io/official_releases/qt/5.12/5.12.0/).

Creating a qt account is part of the process.

**therecipe/qt**

After installing QT, [this document](https://github.com/therecipe/qt/wiki/Installation) and the links at the bottom for your system (linux, windows or macOS) should guide on how to install the go qt mindings.

In the **installation on {system}** links at the bottom, be sure to go through the *Official version (with iOS/Android support)* section to finalize the install

## Setup
### Clone this repo
```bash
git clone git@github.com:acemasterjb/qtips.git
```

### Add the vendor
Run this command before deploying (and maybe before running `qtsetup` command from the [Prerequisites](#Prerequisites) step):
```bash
go mod vendor
```

### Deploy the app
The "app" is defined in main.go for now, to deploy the bindings run:
```bash
qtdeploy build desktop
```

... the `desktop` bit is important, it will deploy an executable app for your system (windows, linux or macOS). You can change this to deploy and app for another platform [if it is supported](https://github.com/therecipe/qt/wiki/Deploying-Application)

### Run the app
For me, my deployed app (linux for linux) was deployed to

```./deploy/linux/qt_eg```

Your app name may be different, just be sure to peep what yours was named before running:
```bash
./deploy/linux/[APP NAME HERE]
```
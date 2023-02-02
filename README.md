# Car game 
## Beschrijving
Spel waarbij je als auto de aankomende auto's moet ontwijken.
Gemaakt door Daniëlle Kwakkel (djkwakkel@che.nl) voor het OpenSource event op 3 februari 2023.

## Installatie
- Download de code in deze repo
- Download de Arduino IDE (van https://www.arduino.cc/en/software of gebruik een package manager als Scoop of Chocolaty als je die hebt)
- Open de GO code en doe een go build
- Sluit de Arduino aan (natuurlijk nadat je de elektronica hebt gebouwd en laten controleren door een docent)
- Doe een go run

## Packages
Het spel bestaat uit 4 packages en een main.go programma. 

- **globalsettings** <br/>
In de globalsettings package worden de globale settings bijgehouden over het spel, zoals de grootte van het scherm, de window, en de initiele snelheid van de auto's. De struct die hiervoor wordt geïnitialiseerd wordt in de andere packages aangeroepen om deze globale variabelen bij te houden. Er wordt gerefereerd aan dit struct door _gs.S._ te gebruiken. 

- **cars** <br/>
Deze package bevat alle variabelen en functies voor de cars. Er zijn twee soorten auto's: inkomende auto's en de speler's auto. De inkomende auto's worden bijgehouden in _Cars ([]*Car)_. De spelers auto is globaal gescoped als _PlayerCar (*Car)_. 

- **visuals** <br/>
De visuals package bevat de functies voor het laden, schalen en weergeven van de afbeeldingen en tekst.

- **serialcon** <br/>
De serialcon package bevat alle functies voor het opzetten en laden van de serial connectie. Er wordt een channel opgezet die de serial input (bijv. verzonden door Arduino) bijhoudt. Door de serial boolean aan te passen in de global settings kan je het spel spelen met de serial input óf met pijltjes toetsen.

## Veel plezier!

**Ben je af? Je kan het spel opnieuw opstarten door op de spatiebalk te drukken** <br/>
**Als je het spel wilt afsluiten typ je in de terminal _ctrl + C_**


# Car game 
## Beschrijving
Spel waarbij je komende auto's moet ontwijken.
Gemaakt door Daniëlle Kwakkel (djkwakkel@che.nl) voor het OpenSource event op 3 februari 2023.

## Installatie

## Packages
Het spel bestaat uit 4 packages en een main.go programma. 

- **globalsettings** <br/>
In de globalsettings package worden de globale settings bijgehouden over het spel, zoals de grootte van het scherm, de window, en de initiele snelheid van de auto's. De struct die hiervoor wordt geïnitialiseerd wordt in de andere packages aangeroepen om deze globale variabelen bij te houden. Er wordt gerefereerd aan dit struct door _gs.S._ te gebruiken. 

- **cars** <br/>
Deze package bevat alle variabelen en functies voor de cars. Er zijn twee soorten auto's: inkomende auto's en de speler's auto. De inkomende auto's worden bijgehouden in _Cars ([]*Car)_. De spelers auto is globaal gescoped als _PlayerCar (*Car)_. 

- **visuals** <br/>
De visuals package bevat de functies voor het laden, schalen en weergeven van de afbeeldingen en tekst.

- **serialcon** <br/>
De serialcon package bevat alle functies voor het opzetten en laden van de serial connectie. Er wordt een channel opgezet die de serial input (bijv. verzonden door Arduino) bijhoudt. 

## Hoe werkt het



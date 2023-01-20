/* Open Source DEV workshop - 3 february 2023
 * Made by Daniëlle Kwakkel
 * Based on the tutorial by Last minute engineers: https://lastminuteengineers.com/rotary-encoder-arduino-tutorial/
 */
 
// Rotary Encoder Inputs
#define CLK 4
#define DT 3
#define SW 2

int counter = 100;
int currentStateCLK;
int lastStateCLK;
String currentDir ="";
unsigned long lastButtonPress = 0;

void setup() {
	
	// Set encoder pins as inputs
	pinMode(CLK,INPUT);
	pinMode(DT,INPUT);
	pinMode(SW, INPUT_PULLUP);

	// Setup Serial Monitor
	Serial.begin(9600);

	// Read the initial state of CLK
	lastStateCLK = digitalRead(CLK);
}

void loop() {
	//Serial.println((String)"CLK" + digitalRead(CLK) +  "DT" +  digitalRead(DT) + "SW" + digitalRead(SW));
	// Read the current state of CLK
	currentStateCLK = digitalRead(CLK);

	// If last and current state of CLK are different, then pulse occurred
	// React to only 1 state change to avoid double count
	if (currentStateCLK != lastStateCLK  && currentStateCLK == 1){

		// If the DT state is different than the CLK state then
		// the encoder is rotating CCW so decrement
		if (digitalRead(DT) != currentStateCLK) {
			counter --;
			currentDir ="CCW";
		} else {
			// Encoder is rotating CW so increment
			counter ++;
			currentDir ="CW";
		}

		//Serial.print("Direction: ");
		//Serial.print(currentDir);
		//Serial.print(" | Counter: ");
		//Serial.println(counter);
    Serial.print("X");
    Serial.println(counter);
	}

	// Remember last CLK state
	lastStateCLK = currentStateCLK;

	// Read the button state
	int btnState = digitalRead(SW);

	//If we detect LOW signal, button is pressed
	if (btnState == LOW) {
		//if 50ms have passed since last LOW pulse, it means that the
		//button has been pressed, released and pressed again
		if (millis() - lastButtonPress > 50) {
			Serial.println("Button pressed!");
		}

		// Remember last button press event
		lastButtonPress = millis();
	}

	// Put in a slight delay to help debounce the reading
	delay(1);
}
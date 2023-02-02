/* Open Source DEV workshop - 3 february 2023
 * Made by Daniëlle Kwakkel
 */

int leftButtonState;
int rightButtonState;
int leftPin = 7;
int rightPin = 8;

int counter = 100;

void setup() {
  Serial.begin(9600);
  pinMode(leftPin, INPUT_PULLUP);
  pinMode(rightPin, INPUT_PULLUP);
}

void loop() {
  leftButtonState = digitalRead(leftPin);
  rightButtonState = digitalRead(rightPin);

  //printOutputs(leftButtonState, rightButtonState); // for debugging

  if ((leftButtonState == 0) && (rightButtonState != 0)){
    if (counter > 0){
    counter -= 1;
    }
  } else if ((leftButtonState != 0) && (rightButtonState == 0)){
    if (counter < 200){    
      counter++;
      }
  }
  
  Serial.print("X");
  Serial.println(counter);
  delay(10);
}

void printOutputs(int left, int right){
  Serial.print("Left: "); Serial.println(left);
  Serial.print("Right: "); Serial.println(right);
}

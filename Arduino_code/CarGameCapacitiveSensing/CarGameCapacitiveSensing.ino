#include <CapacitiveSensor.h>

/* Open Source DEV workshop - 3 february 2023
 * Made by Daniëlle Kwakkel
 * Based on the CapitiveSense Library Demo Sketch by Paul Badger 2008
 * Uses a high value resistor e.g. 10M between send pin and receive pin
 * Receive pin is the sensor pin - try different amounts of foil/metal on this pin
 */


CapacitiveSensor   cs_9_8 = CapacitiveSensor(9,8);        // 10M resistor between pins 9 & 8, pin 8  is sensor pin, add a wire and or foil if desired
CapacitiveSensor   cs_9_10 = CapacitiveSensor(9,10);        // 10M resistor between pins 9 & 10, pin 10 is sensor pin, add a wire and or foil
int counter = 100;

void setup()                    
{
   Serial.begin(9600);
}

void loop()                    
{
    long start = millis();
    long left =  cs_9_8.capacitiveSensor(30);
    long right =  cs_9_10.capacitiveSensor(30);

    //printOutputs(start, left, right); // for debugging
    
    if ((left > 1000) && !(right > 1000)){
      if (counter > 0){
      counter -= 1;
      }
    } else if (!(left > 1000) && (right > 1000)){
      if (counter < 200){
        counter++;
      }
    }
    
    Serial.print("X");
    Serial.println(counter);
    delay(10);                             // delay to limit data to serial port 
}

void printOutputs(long start, long left, long right){
    Serial.print(millis() - start);        // check on performance in milliseconds
    Serial.print("\t");                    // tab character for debug windown spacing
    Serial.print("L");
    Serial.print(left);                  // print sensor output 1
    Serial.print("\t");
    Serial.print("R");
    Serial.println(right);                  // print sensor output 2
}


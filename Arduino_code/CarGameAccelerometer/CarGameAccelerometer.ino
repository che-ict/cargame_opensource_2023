/* Open Source DEV workshop - 3 february 2023
 * Made by Daniëlle Kwakkel
 * Based on code from Noah Shibley (https://github.com/n0m1/MMA8453_n0m1/blob/master/examples/MMA8453_n0m1_dataMode/MMA8453_n0m1_dataMode.ino)
 */

#include <Wire.h>
#include <MMA8453_n0m1.h>

MMA8453_n0m1 accel;
byte counter = 100;

void setup()
{
  Serial.begin(9600);
  accel.setI2CAddr(0x1D); //change your device address if necessary, default is 0x1C
  accel.dataMode(true, 2); //enable highRes 10bit, 2g range [2g,4g,8g]
}

void loop()
{
  accel.update();
  //printOutputs(); // For debugging

    if (accel.y() < -150){
    if (counter > 0){
      counter -= 1;
    }
  } else if (accel.y() > 150){
    if (counter < 200){
      counter++;
    }
  }
  
  Serial.print("X");
  Serial.println(counter);


  delay(10);
}

void printOutputs(){
  Serial.print("x: ");
  Serial.print(accel.x());
  Serial.print(" y: ");
  Serial.print(accel.y());
  Serial.print(" z: ");
  Serial.println(accel.z());
}
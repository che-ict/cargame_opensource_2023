/* Open Source DEV workshop - 3 february 2023
 * Made by Daniëlle Kwakkel
 * Based on code from https://pimylifeup.com/arduino-accelerometer-adxl345/
 */
 
#include <Wire.h>
#include <Adafruit_Sensor.h> 
#include <Adafruit_ADXL345_U.h>

Adafruit_ADXL345_Unified accel = Adafruit_ADXL345_Unified();
int counter = 100;

void setup(void) 
{
  Serial.begin(9600);  
  if(!accel.begin())
  {
    Serial.println("No ADXL345 sensor detected.");
    while(1);
  }
}
void loop(void) 
{
  sensors_event_t event; 
  accel.getEvent(&event);
    
  //printOutputs(event); //for debugging
  if (event.acceleration.y > 2){
    if (counter > 0){
      counter -= 1;
    }
  } else if (event.acceleration.y < -2){
    if (counter < 200){
      counter++;
    }
  }
  
  Serial.print("X");
  Serial.println(counter);

  delay(10);
}

void printOutputs(sensors_event_t event){
  Serial.print("X: "); Serial.print(event.acceleration.x); Serial.print("  ");
  Serial.print("Y: "); Serial.print(event.acceleration.y); Serial.print("  ");
  Serial.print("Z: "); Serial.print(event.acceleration.z); Serial.print("  ");
  Serial.println("m/s^2 ");
}
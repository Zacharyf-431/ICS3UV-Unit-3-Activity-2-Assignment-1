/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-19
 * @fileoverview: This file rounds the given numbers 
 */

// assign constants 
const num1 = 8.5467;
const num2 = 9.6382;
const num3 = 18.5146;
const num4 = 125.496;

// INPUT - No input

// OUTPUT
// rounded values with field widths
console.log(`${num1.toFixed(3).padStart(10)}`);
console.log(`${num2.toFixed(5).padStart(8)}`);
console.log(`${num3.toFixed(1).padStart(6)}`);
console.log(`${num4.toFixed(1).padStart(3)}`);
// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: Sep 2020
'use strict'
/**
 * This function calculates area and perimeter of rectangle.
 */
function myButtonClicked () {
  // input
  const aBase = parseFloat(document.getElementById('a-base').value)
  const bBase = parseFloat(document.getElementById('b-base').value)
  const height = parseFloat(document.getElementById('height').value)


  // process
  const area = (( aBase + bBase ) / 2 ) * height

  // output
  document.getElementById('area').innerHTML = 'The area is: ' + area + ' mm²'
}
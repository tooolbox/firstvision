/**
 * Payment Gateway API Specification
 * Payment Gateway API for payment processing. 
 *
 * OpenAPI spec version: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.FirstApiSdk);
  }
}(this, function(expect, FirstApiSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new FirstApiSdk.SplitShipment();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('SplitShipment', function() {
    it('should create an instance of SplitShipment', function() {
      // uncomment below and update the code to test SplitShipment
      //var instane = new FirstApiSdk.SplitShipment();
      //expect(instance).to.be.a(FirstApiSdk.SplitShipment);
    });

    it('should have the property totalCount (base name: "totalCount")', function() {
      // uncomment below and update the code to test the property totalCount
      //var instane = new FirstApiSdk.SplitShipment();
      //expect(instance).to.be();
    });

    it('should have the property finalShipment (base name: "finalShipment")', function() {
      // uncomment below and update the code to test the property finalShipment
      //var instane = new FirstApiSdk.SplitShipment();
      //expect(instance).to.be();
    });

  });

}));

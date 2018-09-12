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
    instance = new FirstApiSdk.PaymentCardAuthenticationResult();
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

  describe('PaymentCardAuthenticationResult', function() {
    it('should create an instance of PaymentCardAuthenticationResult', function() {
      // uncomment below and update the code to test PaymentCardAuthenticationResult
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be.a(FirstApiSdk.PaymentCardAuthenticationResult);
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be();
    });

    it('should have the property verificationResponse (base name: "verificationResponse")', function() {
      // uncomment below and update the code to test the property verificationResponse
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be();
    });

    it('should have the property payerAuthenticationResponse (base name: "payerAuthenticationResponse")', function() {
      // uncomment below and update the code to test the property payerAuthenticationResponse
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be();
    });

    it('should have the property authenticationValue (base name: "authenticationValue")', function() {
      // uncomment below and update the code to test the property authenticationValue
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be();
    });

    it('should have the property xid (base name: "xid")', function() {
      // uncomment below and update the code to test the property xid
      //var instane = new FirstApiSdk.PaymentCardAuthenticationResult();
      //expect(instance).to.be();
    });

  });

}));

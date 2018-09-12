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


import ApiClient from '../ApiClient';





/**
* The Address model module.
* @module model/Address
* @version 6.3.0
*/
export default class Address {
    /**
    * Constructs a new <code>Address</code>.
    * @alias module:model/Address
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>Address</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/Address} obj Optional instance to populate.
    * @return {module:model/Address} The populated <code>Address</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Address();

            
            
            

            if (data.hasOwnProperty('company')) {
                obj['company'] = ApiClient.convertToType(data['company'], 'String');
            }
            if (data.hasOwnProperty('address1')) {
                obj['address1'] = ApiClient.convertToType(data['address1'], 'String');
            }
            if (data.hasOwnProperty('address2')) {
                obj['address2'] = ApiClient.convertToType(data['address2'], 'String');
            }
            if (data.hasOwnProperty('city')) {
                obj['city'] = ApiClient.convertToType(data['city'], 'String');
            }
            if (data.hasOwnProperty('region')) {
                obj['region'] = ApiClient.convertToType(data['region'], 'String');
            }
            if (data.hasOwnProperty('postalCode')) {
                obj['postalCode'] = ApiClient.convertToType(data['postalCode'], 'String');
            }
            if (data.hasOwnProperty('country')) {
                obj['country'] = ApiClient.convertToType(data['country'], 'String');
            }
        }
        return obj;
    }

    /**
    * If it's a business, enter the name here.
    * @member {String} company
    */
    company = undefined;
    /**
    * street address or P.O.Box number
    * @member {String} address1
    */
    address1 = undefined;
    /**
    * additional address information
    * @member {String} address2
    */
    address2 = undefined;
    /**
    * City or Locality
    * @member {String} city
    */
    city = undefined;
    /**
    * State or Province
    * @member {String} region
    */
    region = undefined;
    /**
    * ZIP code or postal code
    * @member {String} postalCode
    */
    postalCode = undefined;
    /**
    * Country name
    * @member {String} country
    */
    country = undefined;








}



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
* The Expiration model module.
* @module model/Expiration
* @version 6.3.0
*/
export default class Expiration {
    /**
    * Constructs a new <code>Expiration</code>.
    * Required for normal transactions except for payment with &#39;RECURRING&#39; flags
    * @alias module:model/Expiration
    * @class
    * @param month {String} 
    * @param year {String} 
    */

    constructor(month, year) {
        

        
        

        this['month'] = month;this['year'] = year;

        
    }

    /**
    * Constructs a <code>Expiration</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/Expiration} obj Optional instance to populate.
    * @return {module:model/Expiration} The populated <code>Expiration</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Expiration();

            
            
            

            if (data.hasOwnProperty('month')) {
                obj['month'] = ApiClient.convertToType(data['month'], 'String');
            }
            if (data.hasOwnProperty('year')) {
                obj['year'] = ApiClient.convertToType(data['year'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {String} month
    */
    month = undefined;
    /**
    * @member {String} year
    */
    year = undefined;








}



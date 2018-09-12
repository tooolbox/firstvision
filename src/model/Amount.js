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
import AmountComponents from './AmountComponents';





/**
* The Amount model module.
* @module model/Amount
* @version 6.3.0
*/
export default class Amount {
    /**
    * Constructs a new <code>Amount</code>.
    * @alias module:model/Amount
    * @class
    * @param total {Number} 
    * @param currency {String} ISO 4217 currency code (http://en.wikipedia.org/wiki/ISO_4217).
    */

    constructor(total, currency) {
        

        
        

        this['total'] = total;this['currency'] = currency;

        
    }

    /**
    * Constructs a <code>Amount</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/Amount} obj Optional instance to populate.
    * @return {module:model/Amount} The populated <code>Amount</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Amount();

            
            
            

            if (data.hasOwnProperty('total')) {
                obj['total'] = ApiClient.convertToType(data['total'], 'Number');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('components')) {
                obj['components'] = AmountComponents.constructFromObject(data['components']);
            }
        }
        return obj;
    }

    /**
    * @member {Number} total
    */
    total = undefined;
    /**
    * ISO 4217 currency code (http://en.wikipedia.org/wiki/ISO_4217).
    * @member {String} currency
    */
    currency = undefined;
    /**
    * @member {module:model/AmountComponents} components
    */
    components = undefined;








}



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
import PaymentMethod from './PaymentMethod';





/**
* The CardVerificationsTransaction model module.
* @module model/CardVerificationsTransaction
* @version 6.3.0
*/
export default class CardVerificationsTransaction {
    /**
    * Constructs a new <code>CardVerificationsTransaction</code>.
    * Model for card verifcation transactions
    * @alias module:model/CardVerificationsTransaction
    * @class
    * @param paymentMethod {module:model/PaymentMethod} 
    */

    constructor(paymentMethod) {
        

        
        

        this['paymentMethod'] = paymentMethod;

        
    }

    /**
    * Constructs a <code>CardVerificationsTransaction</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/CardVerificationsTransaction} obj Optional instance to populate.
    * @return {module:model/CardVerificationsTransaction} The populated <code>CardVerificationsTransaction</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CardVerificationsTransaction();

            
            
            

            if (data.hasOwnProperty('storeId')) {
                obj['storeId'] = ApiClient.convertToType(data['storeId'], 'String');
            }
            if (data.hasOwnProperty('paymentMethod')) {
                obj['paymentMethod'] = PaymentMethod.constructFromObject(data['paymentMethod']);
            }
        }
        return obj;
    }

    /**
    * Optional Outlet ID for clients that support multiple store in the same app
    * @member {String} storeId
    */
    storeId = undefined;
    /**
    * @member {module:model/PaymentMethod} paymentMethod
    */
    paymentMethod = undefined;








}



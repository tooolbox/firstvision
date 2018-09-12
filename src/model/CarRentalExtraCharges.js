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
* The CarRentalExtraCharges model module.
* @module model/CarRentalExtraCharges
* @version 6.3.0
*/
export default class CarRentalExtraCharges {
    /**
    * Constructs a new <code>CarRentalExtraCharges</code>.
    * @alias module:model/CarRentalExtraCharges
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>CarRentalExtraCharges</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/CarRentalExtraCharges} obj Optional instance to populate.
    * @return {module:model/CarRentalExtraCharges} The populated <code>CarRentalExtraCharges</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CarRentalExtraCharges();

            
            
            

            if (data.hasOwnProperty('chargeItem')) {
                obj['chargeItem'] = ApiClient.convertToType(data['chargeItem'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {module:model/CarRentalExtraCharges.ChargeItemEnum} chargeItem
    */
    chargeItem = undefined;






    /**
    * Allowed values for the <code>chargeItem</code> property.
    * @enum {String}
    * @readonly
    */
    static ChargeItemEnum = {
    
        /**
         * value: "EXTRA_MILEAGE"
         * @const
         */
        "EXTRA_MILEAGE": "EXTRA_MILEAGE",
    
        /**
         * value: "GAS"
         * @const
         */
        "GAS": "GAS",
    
        /**
         * value: "LATE_RETURN"
         * @const
         */
        "LATE_RETURN": "LATE_RETURN",
    
        /**
         * value: "ONE_WAY_SERVICE_FEE"
         * @const
         */
        "ONE_WAY_SERVICE_FEE": "ONE_WAY_SERVICE_FEE",
    
        /**
         * value: "PARKING_VIOLATION"
         * @const
         */
        "PARKING_VIOLATION": "PARKING_VIOLATION"    
    };



}



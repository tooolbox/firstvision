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
* The AirlineTravelRoute model module.
* @module model/AirlineTravelRoute
* @version 6.3.0
*/
export default class AirlineTravelRoute {
    /**
    * Constructs a new <code>AirlineTravelRoute</code>.
    * @alias module:model/AirlineTravelRoute
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>AirlineTravelRoute</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/AirlineTravelRoute} obj Optional instance to populate.
    * @return {module:model/AirlineTravelRoute} The populated <code>AirlineTravelRoute</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AirlineTravelRoute();

            
            
            

            if (data.hasOwnProperty('departureDate')) {
                obj['departureDate'] = ApiClient.convertToType(data['departureDate'], 'Date');
            }
            if (data.hasOwnProperty('origin')) {
                obj['origin'] = ApiClient.convertToType(data['origin'], 'String');
            }
            if (data.hasOwnProperty('destination')) {
                obj['destination'] = ApiClient.convertToType(data['destination'], 'String');
            }
            if (data.hasOwnProperty('carrierCode')) {
                obj['carrierCode'] = ApiClient.convertToType(data['carrierCode'], 'String');
            }
            if (data.hasOwnProperty('serviceClass')) {
                obj['serviceClass'] = ApiClient.convertToType(data['serviceClass'], 'String');
            }
            if (data.hasOwnProperty('stopoverType')) {
                obj['stopoverType'] = ApiClient.convertToType(data['stopoverType'], 'String');
            }
            if (data.hasOwnProperty('fareBasisCode')) {
                obj['fareBasisCode'] = ApiClient.convertToType(data['fareBasisCode'], 'String');
            }
            if (data.hasOwnProperty('departureTax')) {
                obj['departureTax'] = ApiClient.convertToType(data['departureTax'], 'Number');
            }
            if (data.hasOwnProperty('flightNumber')) {
                obj['flightNumber'] = ApiClient.convertToType(data['flightNumber'], 'String');
            }
        }
        return obj;
    }

    /**
    * Date of departure
    * @member {Date} departureDate
    */
    departureDate = undefined;
    /**
    * @member {String} origin
    */
    origin = undefined;
    /**
    * @member {String} destination
    */
    destination = undefined;
    /**
    * @member {String} carrierCode
    */
    carrierCode = undefined;
    /**
    * @member {String} serviceClass
    */
    serviceClass = undefined;
    /**
    * @member {module:model/AirlineTravelRoute.StopoverTypeEnum} stopoverType
    */
    stopoverType = undefined;
    /**
    * @member {String} fareBasisCode
    */
    fareBasisCode = undefined;
    /**
    * @member {Number} departureTax
    */
    departureTax = undefined;
    /**
    * @member {String} flightNumber
    */
    flightNumber = undefined;






    /**
    * Allowed values for the <code>stopoverType</code> property.
    * @enum {String}
    * @readonly
    */
    static StopoverTypeEnum = {
    
        /**
         * value: "DIRECT"
         * @const
         */
        "DIRECT": "DIRECT",
    
        /**
         * value: "STOPOVER"
         * @const
         */
        "STOPOVER": "STOPOVER"    
    };



}



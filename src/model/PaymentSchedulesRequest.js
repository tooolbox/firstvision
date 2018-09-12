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
import Amount from './Amount';
import ClientLocale from './ClientLocale';
import Frequency from './Frequency';
import PaymentMethod from './PaymentMethod';





/**
* The PaymentSchedulesRequest model module.
* @module model/PaymentSchedulesRequest
* @version 6.3.0
*/
export default class PaymentSchedulesRequest {
    /**
    * Constructs a new <code>PaymentSchedulesRequest</code>.
    * Request to create a gateway scheduler.
    * @alias module:model/PaymentSchedulesRequest
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>PaymentSchedulesRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/PaymentSchedulesRequest} obj Optional instance to populate.
    * @return {module:model/PaymentSchedulesRequest} The populated <code>PaymentSchedulesRequest</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PaymentSchedulesRequest();

            
            
            

            if (data.hasOwnProperty('storeId')) {
                obj['storeId'] = ApiClient.convertToType(data['storeId'], 'String');
            }
            if (data.hasOwnProperty('referenceOrderId')) {
                obj['referenceOrderId'] = ApiClient.convertToType(data['referenceOrderId'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('numberOfPayments')) {
                obj['numberOfPayments'] = ApiClient.convertToType(data['numberOfPayments'], 'Number');
            }
            if (data.hasOwnProperty('maximumFailures')) {
                obj['maximumFailures'] = ApiClient.convertToType(data['maximumFailures'], 'Number');
            }
            if (data.hasOwnProperty('invoiceNumber')) {
                obj['invoiceNumber'] = ApiClient.convertToType(data['invoiceNumber'], 'String');
            }
            if (data.hasOwnProperty('poNumber')) {
                obj['poNumber'] = ApiClient.convertToType(data['poNumber'], 'String');
            }
            if (data.hasOwnProperty('transactionOrigin')) {
                obj['transactionOrigin'] = ApiClient.convertToType(data['transactionOrigin'], 'String');
            }
            if (data.hasOwnProperty('dynamicMerchantName')) {
                obj['dynamicMerchantName'] = ApiClient.convertToType(data['dynamicMerchantName'], 'String');
            }
            if (data.hasOwnProperty('frequency')) {
                obj['frequency'] = Frequency.constructFromObject(data['frequency']);
            }
            if (data.hasOwnProperty('paymentMethod')) {
                obj['paymentMethod'] = PaymentMethod.constructFromObject(data['paymentMethod']);
            }
            if (data.hasOwnProperty('amount')) {
                obj['amount'] = Amount.constructFromObject(data['amount']);
            }
            if (data.hasOwnProperty('clientLocale')) {
                obj['clientLocale'] = ClientLocale.constructFromObject(data['clientLocale']);
            }
            if (data.hasOwnProperty('orderId')) {
                obj['orderId'] = ApiClient.convertToType(data['orderId'], 'String');
            }
        }
        return obj;
    }

    /**
    * Store ID number.
    * @member {String} storeId
    */
    storeId = undefined;
    /**
    * Order ID used to create recurring payment from existing transaction.
    * @member {String} referenceOrderId
    */
    referenceOrderId = undefined;
    /**
    * Date of mandate signature.
    * @member {Date} startDate
    */
    startDate = undefined;
    /**
    * Number of times the recurring payment will process.
    * @member {Number} numberOfPayments
    */
    numberOfPayments = undefined;
    /**
    * Number of failures that can be encountered before re-tries cease.
    * @member {Number} maximumFailures
    */
    maximumFailures = undefined;
    /**
    * Invoice number.
    * @member {String} invoiceNumber
    */
    invoiceNumber = undefined;
    /**
    * Purchase order number.
    * @member {String} poNumber
    */
    poNumber = undefined;
    /**
    * The source of the transaction. The possible values are ECI (if the order was received via email or Internet), MOTO (mail order / telephone order) and RETAIL (face to face).
    * @member {module:model/PaymentSchedulesRequest.TransactionOriginEnum} transactionOrigin
    */
    transactionOrigin = undefined;
    /**
    * Dynamic merchant name for the cardholder‘s statement.
    * @member {String} dynamicMerchantName
    */
    dynamicMerchantName = undefined;
    /**
    * @member {module:model/Frequency} frequency
    */
    frequency = undefined;
    /**
    * @member {module:model/PaymentMethod} paymentMethod
    */
    paymentMethod = undefined;
    /**
    * @member {module:model/Amount} amount
    */
    amount = undefined;
    /**
    * @member {module:model/ClientLocale} clientLocale
    */
    clientLocale = undefined;
    /**
    * Client Order ID if supplied by client.
    * @member {String} orderId
    */
    orderId = undefined;






    /**
    * Allowed values for the <code>transactionOrigin</code> property.
    * @enum {String}
    * @readonly
    */
    static TransactionOriginEnum = {
    
        /**
         * value: "ECI"
         * @const
         */
        "ECI": "ECI",
    
        /**
         * value: "MOTO"
         * @const
         */
        "MOTO": "MOTO",
    
        /**
         * value: "RETAIL"
         * @const
         */
        "RETAIL": "RETAIL"    
    };



}



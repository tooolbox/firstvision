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
import ProcessorData from './ProcessorData';
import ResponseType from './ResponseType';
import TransactionResponseAuthenticationRedirect from './TransactionResponseAuthenticationRedirect';
import TransactionType from './TransactionType';





/**
* The TransactionResponse model module.
* @module model/TransactionResponse
* @version 6.3.0
*/
export default class TransactionResponse {
    /**
    * Constructs a new <code>TransactionResponse</code>.
    * This is the common object for primary and secondary payment request responses
    * @alias module:model/TransactionResponse
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>TransactionResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/TransactionResponse} obj Optional instance to populate.
    * @return {module:model/TransactionResponse} The populated <code>TransactionResponse</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TransactionResponse();

            
            
            

            if (data.hasOwnProperty('responseType')) {
                obj['responseType'] = ResponseType.constructFromObject(data['responseType']);
            }
            if (data.hasOwnProperty('clientRequestId')) {
                obj['clientRequestId'] = ApiClient.convertToType(data['clientRequestId'], 'String');
            }
            if (data.hasOwnProperty('apiTraceId')) {
                obj['apiTraceId'] = ApiClient.convertToType(data['apiTraceId'], 'String');
            }
            if (data.hasOwnProperty('ipgTransactionId')) {
                obj['ipgTransactionId'] = ApiClient.convertToType(data['ipgTransactionId'], 'String');
            }
            if (data.hasOwnProperty('orderId')) {
                obj['orderId'] = ApiClient.convertToType(data['orderId'], 'String');
            }
            if (data.hasOwnProperty('transactionType')) {
                obj['transactionType'] = TransactionType.constructFromObject(data['transactionType']);
            }
            if (data.hasOwnProperty('authorizationCode')) {
                obj['authorizationCode'] = ApiClient.convertToType(data['authorizationCode'], 'String');
            }
            if (data.hasOwnProperty('avsResponse')) {
                obj['avsResponse'] = ApiClient.convertToType(data['avsResponse'], 'String');
            }
            if (data.hasOwnProperty('securityCodeResponse')) {
                obj['securityCodeResponse'] = ApiClient.convertToType(data['securityCodeResponse'], 'String');
            }
            if (data.hasOwnProperty('brand')) {
                obj['brand'] = ApiClient.convertToType(data['brand'], 'String');
            }
            if (data.hasOwnProperty('country')) {
                obj['country'] = ApiClient.convertToType(data['country'], 'String');
            }
            if (data.hasOwnProperty('terminalId')) {
                obj['terminalId'] = ApiClient.convertToType(data['terminalId'], 'String');
            }
            if (data.hasOwnProperty('clientTransactionId')) {
                obj['clientTransactionId'] = ApiClient.convertToType(data['clientTransactionId'], 'String');
            }
            if (data.hasOwnProperty('transactionTime')) {
                obj['transactionTime'] = ApiClient.convertToType(data['transactionTime'], 'Number');
            }
            if (data.hasOwnProperty('approvedAmount')) {
                obj['approvedAmount'] = Amount.constructFromObject(data['approvedAmount']);
            }
            if (data.hasOwnProperty('transactionStatus')) {
                obj['transactionStatus'] = ApiClient.convertToType(data['transactionStatus'], 'String');
            }
            if (data.hasOwnProperty('transactionState')) {
                obj['transactionState'] = ApiClient.convertToType(data['transactionState'], 'String');
            }
            if (data.hasOwnProperty('authenticationRedirect')) {
                obj['authenticationRedirect'] = TransactionResponseAuthenticationRedirect.constructFromObject(data['authenticationRedirect']);
            }
            if (data.hasOwnProperty('schemeTransactionId')) {
                obj['schemeTransactionId'] = ApiClient.convertToType(data['schemeTransactionId'], 'String');
            }
            if (data.hasOwnProperty('processor')) {
                obj['processor'] = ProcessorData.constructFromObject(data['processor']);
            }
        }
        return obj;
    }

    /**
    * The schema type returned in the response.
    * @member {module:model/ResponseType} responseType
    */
    responseType = undefined;
    /**
    * Echoes back the value in the Request header
    * @member {String} clientRequestId
    */
    clientRequestId = undefined;
    /**
    * Echoes back the value in the Request header
    * @member {String} apiTraceId
    */
    apiTraceId = undefined;
    /**
    * The Response Transaction ID
    * @member {String} ipgTransactionId
    */
    ipgTransactionId = undefined;
    /**
    * Client Order ID if supplied by client, otherwise the Order ID
    * @member {String} orderId
    */
    orderId = undefined;
    /**
    * @member {module:model/TransactionType} transactionType
    */
    transactionType = undefined;
    /**
    * The processor approval code for compliance.
    * @member {String} authorizationCode
    */
    authorizationCode = undefined;
    /**
    * The processor address validation response for compliance.
    * @member {String} avsResponse
    */
    avsResponse = undefined;
    /**
    * The processor card verification validation response for compliance.
    * @member {String} securityCodeResponse
    */
    securityCodeResponse = undefined;
    /**
    * Card brand of the payment instrument
    * @member {String} brand
    */
    brand = undefined;
    /**
    * Country of the card issuer
    * @member {String} country
    */
    country = undefined;
    /**
    * The terminal that is processing the transaction
    * @member {String} terminalId
    */
    terminalId = undefined;
    /**
    * The unique client Transaction ID from the Request header, if supplied
    * @member {String} clientTransactionId
    */
    clientTransactionId = undefined;
    /**
    * The transaction time in seconds since Epoch
    * @member {Number} transactionTime
    */
    transactionTime = undefined;
    /**
    * @member {module:model/Amount} approvedAmount
    */
    approvedAmount = undefined;
    /**
    * The status of the transaction. APPROVED/WAITING are returned by the endpoints.  VALIDATION_FAILED/DECLINED are errors. See ErrorResponse object for details.
    * @member {module:model/TransactionResponse.TransactionStatusEnum} transactionStatus
    */
    transactionStatus = undefined;
    /**
    * The state of the transaction.
    * @member {module:model/TransactionResponse.TransactionStateEnum} transactionState
    */
    transactionState = undefined;
    /**
    * @member {module:model/TransactionResponseAuthenticationRedirect} authenticationRedirect
    */
    authenticationRedirect = undefined;
    /**
    * The transaction id received from schemes for the initial transaction, returned for the transactions marked as \"FIRST\"
    * @member {String} schemeTransactionId
    */
    schemeTransactionId = undefined;
    /**
    * @member {module:model/ProcessorData} processor
    */
    processor = undefined;






    /**
    * Allowed values for the <code>transactionStatus</code> property.
    * @enum {String}
    * @readonly
    */
    static TransactionStatusEnum = {
    
        /**
         * value: "APPROVED"
         * @const
         */
        "APPROVED": "APPROVED",
    
        /**
         * value: "WAITING"
         * @const
         */
        "WAITING": "WAITING",
    
        /**
         * value: "VALIDATION_FAILED"
         * @const
         */
        "VALIDATION_FAILED": "VALIDATION_FAILED",
    
        /**
         * value: "DECLINED"
         * @const
         */
        "DECLINED": "DECLINED"    
    };

    /**
    * Allowed values for the <code>transactionState</code> property.
    * @enum {String}
    * @readonly
    */
    static TransactionStateEnum = {
    
        /**
         * value: "AUTHORIZED"
         * @const
         */
        "AUTHORIZED": "AUTHORIZED",
    
        /**
         * value: "CAPTURED"
         * @const
         */
        "CAPTURED": "CAPTURED",
    
        /**
         * value: "COMPLETED_GET"
         * @const
         */
        "COMPLETED_GET": "COMPLETED_GET",
    
        /**
         * value: "DECLINED"
         * @const
         */
        "DECLINED": "DECLINED",
    
        /**
         * value: "CHECKED"
         * @const
         */
        "CHECKED": "CHECKED",
    
        /**
         * value: "INITIALIZED"
         * @const
         */
        "INITIALIZED": "INITIALIZED",
    
        /**
         * value: "PENDING_AUTHORIZATION"
         * @const
         */
        "PENDING_AUTHORIZATION": "PENDING_AUTHORIZATION",
    
        /**
         * value: "PENDING_AUTOVOID"
         * @const
         */
        "PENDING_AUTOVOID": "PENDING_AUTOVOID",
    
        /**
         * value: "PENDING_CAPTURE"
         * @const
         */
        "PENDING_CAPTURE": "PENDING_CAPTURE",
    
        /**
         * value: "PENDING_CREDIT"
         * @const
         */
        "PENDING_CREDIT": "PENDING_CREDIT",
    
        /**
         * value: "PENDING_GIROPAY_INIT"
         * @const
         */
        "PENDING_GIROPAY_INIT": "PENDING_GIROPAY_INIT",
    
        /**
         * value: "PENDING_IDEAL_INIT"
         * @const
         */
        "PENDING_IDEAL_INIT": "PENDING_IDEAL_INIT",
    
        /**
         * value: "PENDING_INIT"
         * @const
         */
        "PENDING_INIT": "PENDING_INIT",
    
        /**
         * value: "PENDING_READY"
         * @const
         */
        "PENDING_READY": "PENDING_READY",
    
        /**
         * value: "PENDING_RETURN"
         * @const
         */
        "PENDING_RETURN": "PENDING_RETURN",
    
        /**
         * value: "PENDING_SETTLEMENT"
         * @const
         */
        "PENDING_SETTLEMENT": "PENDING_SETTLEMENT",
    
        /**
         * value: "PENDING_SOFORT_INIT"
         * @const
         */
        "PENDING_SOFORT_INIT": "PENDING_SOFORT_INIT",
    
        /**
         * value: "PENDING_VOID"
         * @const
         */
        "PENDING_VOID": "PENDING_VOID",
    
        /**
         * value: "READY"
         * @const
         */
        "READY": "READY",
    
        /**
         * value: "SETTLED"
         * @const
         */
        "SETTLED": "SETTLED",
    
        /**
         * value: "VOIDED"
         * @const
         */
        "VOIDED": "VOIDED",
    
        /**
         * value: "WAITING"
         * @const
         */
        "WAITING": "WAITING",
    
        /**
         * value: "WAITING_AUTHENTICATION"
         * @const
         */
        "WAITING_AUTHENTICATION": "WAITING_AUTHENTICATION",
    
        /**
         * value: "WAITING_3D_SECURE"
         * @const
         */
        "WAITING_3D_SECURE": "WAITING_3D_SECURE",
    
        /**
         * value: "WAITING_CLICK_AND_BUY"
         * @const
         */
        "WAITING_CLICK_AND_BUY": "WAITING_CLICK_AND_BUY",
    
        /**
         * value: "WAITING_GIROPAY"
         * @const
         */
        "WAITING_GIROPAY": "WAITING_GIROPAY",
    
        /**
         * value: "WAITING_IDEAL"
         * @const
         */
        "WAITING_IDEAL": "WAITING_IDEAL",
    
        /**
         * value: "WAITING_KLARNA"
         * @const
         */
        "WAITING_KLARNA": "WAITING_KLARNA",
    
        /**
         * value: "WAITING_PAYPAL"
         * @const
         */
        "WAITING_PAYPAL": "WAITING_PAYPAL",
    
        /**
         * value: "WAITING_PAYPAL_EVENT"
         * @const
         */
        "WAITING_PAYPAL_EVENT": "WAITING_PAYPAL_EVENT",
    
        /**
         * value: "WAITING_PPRO_LONG_WAITING"
         * @const
         */
        "WAITING_PPRO_LONG_WAITING": "WAITING_PPRO_LONG_WAITING",
    
        /**
         * value: "WAITING_SOFORT"
         * @const
         */
        "WAITING_SOFORT": "WAITING_SOFORT",
    
        /**
         * value: "WAITING_T_PAY"
         * @const
         */
        "WAITING_T_PAY": "WAITING_T_PAY",
    
        /**
         * value: "WAITING_ALIPAY_PAYSECURE"
         * @const
         */
        "WAITING_ALIPAY_PAYSECURE": "WAITING_ALIPAY_PAYSECURE"    
    };



}



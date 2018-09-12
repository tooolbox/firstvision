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


import ApiClient from "../ApiClient";
import ErrorResponse from '../model/ErrorResponse';
import SecondaryTransaction from '../model/SecondaryTransaction';
import TransactionResponse from '../model/TransactionResponse';

/**
* Order service.
* @module api/OrderApi
* @version 6.3.0
*/
export default class OrderApi {

    /**
    * Constructs a new OrderApi. 
    * @alias module:api/OrderApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Use this to capture/complete a transaction. Partial postauths are allowed.
     * This can be used for postauth and partial postauths.
     * @param {module:model/String} contentType content type
     * @param {String} clientRequestId A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
     * @param {String} apiKey 
     * @param {Number} timestamp Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
     * @param {String} messageSignature Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256  algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
     * @param {String} orderId Gateway order identifier as returned in the parameter orderId
     * @param {module:model/SecondaryTransaction} payload 
     * @param {Object} opts Optional parameters
     * @param {String} opts.storeId an optional outlet id for clients that support multiple store in the same developer app
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TransactionResponse} and HTTP response
     */
    performPaymentPostAuthorisationWithHttpInfo(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts) {
      opts = opts || {};
      let postBody = payload;

      // verify the required parameter 'contentType' is set
      if (contentType === undefined || contentType === null) {
        throw new Error("Missing the required parameter 'contentType' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'clientRequestId' is set
      if (clientRequestId === undefined || clientRequestId === null) {
        throw new Error("Missing the required parameter 'clientRequestId' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'apiKey' is set
      if (apiKey === undefined || apiKey === null) {
        throw new Error("Missing the required parameter 'apiKey' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'timestamp' is set
      if (timestamp === undefined || timestamp === null) {
        throw new Error("Missing the required parameter 'timestamp' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'messageSignature' is set
      if (messageSignature === undefined || messageSignature === null) {
        throw new Error("Missing the required parameter 'messageSignature' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'orderId' is set
      if (orderId === undefined || orderId === null) {
        throw new Error("Missing the required parameter 'orderId' when calling performPaymentPostAuthorisation");
      }

      // verify the required parameter 'payload' is set
      if (payload === undefined || payload === null) {
        throw new Error("Missing the required parameter 'payload' when calling performPaymentPostAuthorisation");
      }


      let pathParams = {
        'order-id': orderId
      };
      let queryParams = {
        'storeId': opts['storeId']
      };
      let headerParams = {
        'Content-Type': contentType,
        'Client-Request-Id': clientRequestId,
        'Api-Key': apiKey,
        'Timestamp': timestamp,
        'Message-Signature': messageSignature
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = TransactionResponse;

      return this.apiClient.callApi(
        '/v1/orders/{order-id}/postauth', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Use this to capture/complete a transaction. Partial postauths are allowed.
     * This can be used for postauth and partial postauths.
     * @param {module:model/String} contentType content type
     * @param {String} clientRequestId A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
     * @param {String} apiKey 
     * @param {Number} timestamp Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
     * @param {String} messageSignature Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256  algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
     * @param {String} orderId Gateway order identifier as returned in the parameter orderId
     * @param {module:model/SecondaryTransaction} payload 
     * @param {Object} opts Optional parameters
     * @param {String} opts.storeId an optional outlet id for clients that support multiple store in the same developer app
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TransactionResponse}
     */
    performPaymentPostAuthorisation(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts) {
      return this.performPaymentPostAuthorisationWithHttpInfo(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Use this to return/refund on the order. Partial returns are allowed.
     * This can be used for Returns and Partial Returns.
     * @param {module:model/String} contentType content type
     * @param {String} clientRequestId A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
     * @param {String} apiKey 
     * @param {Number} timestamp Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
     * @param {String} messageSignature Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256  algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
     * @param {String} orderId Gateway order identifier as returned in the parameter orderId
     * @param {module:model/SecondaryTransaction} payload 
     * @param {Object} opts Optional parameters
     * @param {String} opts.storeId an optional outlet id for clients that support multiple store in the same developer app
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TransactionResponse} and HTTP response
     */
    returnTransactionWithHttpInfo(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts) {
      opts = opts || {};
      let postBody = payload;

      // verify the required parameter 'contentType' is set
      if (contentType === undefined || contentType === null) {
        throw new Error("Missing the required parameter 'contentType' when calling returnTransaction");
      }

      // verify the required parameter 'clientRequestId' is set
      if (clientRequestId === undefined || clientRequestId === null) {
        throw new Error("Missing the required parameter 'clientRequestId' when calling returnTransaction");
      }

      // verify the required parameter 'apiKey' is set
      if (apiKey === undefined || apiKey === null) {
        throw new Error("Missing the required parameter 'apiKey' when calling returnTransaction");
      }

      // verify the required parameter 'timestamp' is set
      if (timestamp === undefined || timestamp === null) {
        throw new Error("Missing the required parameter 'timestamp' when calling returnTransaction");
      }

      // verify the required parameter 'messageSignature' is set
      if (messageSignature === undefined || messageSignature === null) {
        throw new Error("Missing the required parameter 'messageSignature' when calling returnTransaction");
      }

      // verify the required parameter 'orderId' is set
      if (orderId === undefined || orderId === null) {
        throw new Error("Missing the required parameter 'orderId' when calling returnTransaction");
      }

      // verify the required parameter 'payload' is set
      if (payload === undefined || payload === null) {
        throw new Error("Missing the required parameter 'payload' when calling returnTransaction");
      }


      let pathParams = {
        'order-id': orderId
      };
      let queryParams = {
        'storeId': opts['storeId']
      };
      let headerParams = {
        'Content-Type': contentType,
        'Client-Request-Id': clientRequestId,
        'Api-Key': apiKey,
        'Timestamp': timestamp,
        'Message-Signature': messageSignature
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = TransactionResponse;

      return this.apiClient.callApi(
        '/v1/orders/{order-id}/return', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Use this to return/refund on the order. Partial returns are allowed.
     * This can be used for Returns and Partial Returns.
     * @param {module:model/String} contentType content type
     * @param {String} clientRequestId A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
     * @param {String} apiKey 
     * @param {Number} timestamp Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
     * @param {String} messageSignature Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256  algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
     * @param {String} orderId Gateway order identifier as returned in the parameter orderId
     * @param {module:model/SecondaryTransaction} payload 
     * @param {Object} opts Optional parameters
     * @param {String} opts.storeId an optional outlet id for clients that support multiple store in the same developer app
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TransactionResponse}
     */
    returnTransaction(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts) {
      return this.returnTransactionWithHttpInfo(contentType, clientRequestId, apiKey, timestamp, messageSignature, orderId, payload, opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}

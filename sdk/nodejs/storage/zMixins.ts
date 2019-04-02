// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { Bucket } from "./bucket";
import * as cloudfunctions from "../cloudfunctions";

/**
 * Arguments that can be provided to control the Cloud Function created as the serverless endpoint
 * for a bucket event.
 */
export interface BucketEventCallbackFunctionArgs extends cloudfunctions.CallbackFunctionArgs {
    callback?: BucketEventHandler;
    callbackFactory?: () => BucketEventHandler;

    // never provided.  type them as 'never' so TypeScript will complain if a user tries to include
    // these.

    httpsTriggerUrl?: never;
    triggerHttp?: never;
    eventTrigger?: never;
}

/**
 * Arguments to control how GCP will respond if the Cloud Function fails.  Currently, the only
 * specialized behavior supported is to attempt retrying the Cloud Function.
 * See [cloudfunctions.FailurePolicy] for more information on this.
 */
export interface SimpleBucketEventArgs {
    failurePolicy?: cloudfunctions.FailurePolicy;
}

export interface BucketEventArgs {
    triggerType: "finalize" | "delete" | "archive" | "metadataUpdate";
    failurePolicy?: cloudfunctions.FailurePolicy;
}

/**
 * Shape of the [context] object passed to a Cloud Function when a bucket event fires.
 */
export interface BucketContext extends cloudfunctions.Context {
    /** The type of the event. */
    eventType: "google.storage.object.finalize" | "google.storage.object.delete" | "google.storage.object.archive" | "google.storage.object.metadataUpdate";

    /** The resource that emitted the event. */
    resource: {
        service: "storage.googleapis.com",
        name: string,
        type: "storage#object",
    };
}

/**
 * Shape of the [data] object passed to a Cloud Function when a bucket event fires.
 *
 * See https://cloud.google.com/storage/docs/json_api/v1/objects for more details.
 */
export interface BucketData {
    "kind": "storage#object",

    "bucket": string,
    "contentType": string,
    "crc32c": string,
    "etag": string,
    "generation": number,
    "id": string,
    "md5Hash": string,
    "mediaLink": string,
    "metadata": Record<string, string>;
    "metageneration": number,
    "name": string,
    "selfLink": string,
    "size": number,
    "storageClass": string,
    "timeCreated": string,
    "timeStorageClassUpdated": string,
    "updated": string,
}

export type BucketEventHandler = cloudfunctions.Callback<BucketData, BucketContext, void>;

declare module "./bucket" {
    interface Bucket {
        /**
         * Creates and publishes a Cloud Functions that will be triggered when a new object is
         * created (or an existing object is overwritten, and a new generation of that object is
         * created) in this bucket.
         */
        onObjectFinalized(name: string, handler: BucketEventHandler | BucketEventCallbackFunctionArgs, args?: SimpleBucketEventArgs, opts?: pulumi.ComponentResourceOptions): cloudfunctions.CallbackFunction;

        /**
         * Creates and publishes a Cloud Functions that will be triggered when an object is
         * permanently deleted. Depending on the object versioning setting for a bucket this means:
         *
         * 1. For versioning buckets, this is only sent when a version is permanently deleted (but
         *    not when an object is archived).
         *
         * 2. For non-versioning buckets, this is sent when an object is deleted or overwritten.
         *
         * See https://cloud.google.com/storage/docs/object-versioning for more details.
         */
        onObjectDeleted(name: string, handler: BucketEventHandler | BucketEventCallbackFunctionArgs, args?: SimpleBucketEventArgs, opts?: pulumi.ComponentResourceOptions): cloudfunctions.CallbackFunction;

        /**
         * Creates and publishes a Cloud Functions that will be triggered when a live version of an
         * object is archived or deleted.
         *
         * This event is only sent for versioning buckets.
         *
         * See https://cloud.google.com/storage/docs/object-versioning for more details.
         */
        onObjectArchived(name: string, handler: BucketEventHandler | BucketEventCallbackFunctionArgs, args?: SimpleBucketEventArgs, opts?: pulumi.ComponentResourceOptions): cloudfunctions.CallbackFunction;

        /**
         * Creates and publishes a Cloud Functions that will be triggered when the metadata of an
         * existing object changes.
         *
         * See https://cloud.google.com/storage/docs/metadata for more details.
         */
        onObjectMetadataUpdated(name: string, handler: BucketEventHandler | BucketEventCallbackFunctionArgs, args?: SimpleBucketEventArgs, opts?: pulumi.ComponentResourceOptions): cloudfunctions.CallbackFunction;

        /**
         * Generic helper for registering for any event.
         */
        onObjectEvent(name: string, handler: BucketEventHandler | BucketEventCallbackFunctionArgs, args: BucketEventArgs, opts?: pulumi.ComponentResourceOptions): cloudfunctions.CallbackFunction;
    }
}

Bucket.prototype.onObjectFinalized = function (this: Bucket, name, handler, args, opts) {
    return this.onObjectEvent(name, handler, { ...args, triggerType: "finalize" }, opts);
}

Bucket.prototype.onObjectDeleted = function (this: Bucket, name, handler, args, opts) {
    return this.onObjectEvent(name, handler, { ...args, triggerType: "delete" }, opts);
}

Bucket.prototype.onObjectArchived = function (this: Bucket, name, handler, args, opts) {
    return this.onObjectEvent(name, handler, { ...args, triggerType: "archive" }, opts);
}

Bucket.prototype.onObjectMetadataUpdated = function (this: Bucket, name, handler, args, opts) {
    return this.onObjectEvent(name, handler, { ...args, triggerType: "metadataUpdate" }, opts);
}

Bucket.prototype.onObjectEvent = function (this: Bucket, name, handlerOrCallbackArgs, args, opts = {}) {
    const callbackArgs = handlerOrCallbackArgs instanceof Function
        ? { callback: handlerOrCallbackArgs }
        : handlerOrCallbackArgs;

    return new cloudfunctions.CallbackFunction(name, {
        ...callbackArgs,
        eventTrigger: {
            resource: this.name,
            failurePolicy: args.failurePolicy,
            eventType: `google.storage.object.${args.triggerType}`,
        }
    }, { parent: this, ...opts })
}

// Needed so that typescript won't elide this module
/** @internal */
export const __unused = 0;

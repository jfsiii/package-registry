declare namespace EPR {
    /**
     * Category
     */
    export interface Category {
        id?: string;
        title?: string;
        count?: number;
    }
    /**
     * Dataset
     */
    export interface Dataset {
        title: string;
        name: string;
        release: string;
        type: string;
        ingest_pipeline?: string;
        vars: {
        }[];
        package: string;
    }
    /**
     * Image
     */
    export interface Image {
        src?: string;
        title?: string;
        size?: string;
        type?: string;
    }
    /**
     * Kibana
     */
    export interface Kibana {
        versions?: string;
    }
    /**
     * Package
     */
    export interface Package {
        name: string;
        title?: string;
        version: string;
        readme?: string;
        description: string;
        categories: string[];
        requirement: EPR.Requirement;
        screenshots?: EPR.Image[];
        icons?: EPR.Image[];
        assets?: string[];
        internal?: boolean;
        formatversion: string;
        datasets?: EPR.Dataset[];
        download: string;
        path: string;
    }
    namespace Parameters {
        export type Category = string;
        export type Internal = string;
        export type Kibana = string;
        export type Package = string;
    }
    export interface QueryParameters {
        kibana?: EPR.Parameters.Kibana;
        category?: EPR.Parameters.Category;
        package?: EPR.Parameters.Package;
        internal?: EPR.Parameters.Internal;
    }
    /**
     * Requirement
     */
    export interface Requirement {
        Kibana: EPR.Kibana;
    }
    namespace Responses {
        export type $200 = EPR.Package;
    }
    /**
     * Root-Info
     */
    export interface RootInfo {
        version?: string;
        "service.name"?: string;
    }
}

export namespace backend {
	
	export class ConfigSummary {
	    language: string;
	    logLevel: string;
	    logDir: string;
	
	    static createFrom(source: any = {}) {
	        return new ConfigSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.language = source["language"];
	        this.logLevel = source["logLevel"];
	        this.logDir = source["logDir"];
	    }
	}
	export class LogWriteResult {
	    ok: boolean;
	    error?: string;
	    logDir?: string;
	
	    static createFrom(source: any = {}) {
	        return new LogWriteResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	        this.logDir = source["logDir"];
	    }
	}

}

export namespace systeminfo {
	
	export class Info {
	    version: string;
	    buildTime: string;
	    environment: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.buildTime = source["buildTime"];
	        this.environment = source["environment"];
	    }
	}

}


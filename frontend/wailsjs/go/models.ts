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
	export class UpdateCheckResponse {
	    info: updater.UpdateInfo;
	    currentVersion: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCheckResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.info = this.convertValues(source["info"], updater.UpdateInfo);
	        this.currentVersion = source["currentVersion"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
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

export namespace updater {
	
	export class InstallResult {
	    Path: string;
	
	    static createFrom(source: any = {}) {
	        return new InstallResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	    }
	}
	export class UpdateInfo {
	    protocolVersion: string;
	    latestVersion: string;
	    force: boolean;
	    channel: string;
	    downloadUrl: string;
	    checksum: string;
	    releaseNotes: string;
	    minSupportedVersion: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.protocolVersion = source["protocolVersion"];
	        this.latestVersion = source["latestVersion"];
	        this.force = source["force"];
	        this.channel = source["channel"];
	        this.downloadUrl = source["downloadUrl"];
	        this.checksum = source["checksum"];
	        this.releaseNotes = source["releaseNotes"];
	        this.minSupportedVersion = source["minSupportedVersion"];
	    }
	}

}


"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var default_retrieve_file_handler_1 = require("./default-retrieve-file-handler");
var sourceMapSupport = require("source-map-support");
var transpile_if_ts_1 = require("./transpile-if-ts");
exports.transpileIfTypescript = transpile_if_ts_1.transpileIfTypescript;
function install() {
    var options = {};
    options.retrieveFile = default_retrieve_file_handler_1.defaultRetrieveFileHandler;
    options.emptyCacheBetweenOperations = true;
    options['environment'] = 'node';
    return sourceMapSupport.install(options);
}
exports.install = install;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5kZXguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvaW5kZXgudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7QUFBQSxpRkFBNkU7QUFDN0UscURBQXVEO0FBRXZELHFEQUEwRDtBQUFqRCxrREFBQSxxQkFBcUIsQ0FBQTtBQUM5QjtJQUNFLElBQUksT0FBTyxHQUE2QixFQUFFLENBQUM7SUFDM0MsT0FBTyxDQUFDLFlBQVksR0FBRywwREFBMEIsQ0FBQztJQUNsRCxPQUFPLENBQUMsMkJBQTJCLEdBQUcsSUFBSSxDQUFDO0lBQzNDLE9BQU8sQ0FBQyxhQUFhLENBQUMsR0FBRyxNQUFNLENBQUM7SUFFaEMsTUFBTSxDQUFDLGdCQUFnQixDQUFDLE9BQU8sQ0FBQyxPQUFPLENBQUMsQ0FBQztBQUMzQyxDQUFDO0FBUEQsMEJBT0MifQ==
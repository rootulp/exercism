"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var tsc = require("typescript");
var utils_1 = require("./utils");
function transpileIfTypescript(path, contents, config) {
    if (path && (path.endsWith('.tsx') || path.endsWith('.ts'))) {
        var transpiled = tsc.transpileModule(contents, {
            compilerOptions: utils_1.getTSConfig(config || { __TS_CONFIG__: global['__TS_CONFIG__'] }, true),
            fileName: path
        });
        return transpiled.outputText;
    }
    return contents;
}
exports.transpileIfTypescript = transpileIfTypescript;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidHJhbnNwaWxlLWlmLXRzLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3RyYW5zcGlsZS1pZi10cy50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUFBLGdDQUFrQztBQUNsQyxpQ0FBc0M7QUFFdEMsK0JBQXNDLElBQUksRUFBRSxRQUFRLEVBQUUsTUFBTztJQUMzRCxFQUFFLENBQUMsQ0FBQyxJQUFJLElBQUksQ0FBQyxJQUFJLENBQUMsUUFBUSxDQUFDLE1BQU0sQ0FBQyxJQUFJLElBQUksQ0FBQyxRQUFRLENBQUMsS0FBSyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUM7UUFFNUQsSUFBSSxVQUFVLEdBQUcsR0FBRyxDQUFDLGVBQWUsQ0FBQyxRQUFRLEVBQUU7WUFDN0MsZUFBZSxFQUFFLG1CQUFXLENBQUMsTUFBTSxJQUFJLEVBQUUsYUFBYSxFQUFFLE1BQU0sQ0FBQyxlQUFlLENBQUMsRUFBRSxFQUFFLElBQUksQ0FBQztZQUN4RixRQUFRLEVBQUUsSUFBSTtTQUNmLENBQUMsQ0FBQztRQUVILE1BQU0sQ0FBQyxVQUFVLENBQUMsVUFBVSxDQUFDO0lBQy9CLENBQUM7SUFDRCxNQUFNLENBQUMsUUFBUSxDQUFDO0FBQ2xCLENBQUM7QUFYRCxzREFXQyJ9
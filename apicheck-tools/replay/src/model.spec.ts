import {validateRequest} from "./model";

import { expect } from "chai";
import 'mocha';

describe("validate Request", () => {
    it("should return a string error if empty request", () => {
        let emptyRequest = {};
        let response = validateRequest(emptyRequest);
        expect(response).to.not.be.a('null');
        expect(response).to.be.a('string');
    });

    it("should return a null if valid request", () => {
        let validRequest = {
            "request":{
                "url": "http://example.com/v1/status"
            }
        };

        let response = validateRequest(validRequest);
        expect(response).to.be.a('null');
    });
});

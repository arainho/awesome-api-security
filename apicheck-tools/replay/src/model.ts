import { Validator } from 'jsonschema';

const requestValidator = new Validator();

const requestSchema = {
  "type": "object",
  "properties": {
      "_meta": {
        "type": "object"
      },
      "request": {
        "type": "object",
        "properties": {
          "url": {
            "type": "string",
            "pattern": "https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}"
          },
          "method": {
            "type": "string"
          },
          "headers": {
            "type": "object"
          },
          "body": {
            "type": "string"
          }
        },
        "required": ["url"]
      }
  },
  "required": ["request"]
};

export const validateRequest = (req: any) => {
    let valid = requestValidator.validate(req, requestSchema);
    if (valid.valid) {
        return null;
    }else{
        let errors = valid.errors.join("\n");
        return errors;
    }
}

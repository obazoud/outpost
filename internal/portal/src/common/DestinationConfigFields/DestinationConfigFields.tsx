import { useState, useRef, useEffect } from "react";
import {
  Destination,
  DestinationTypeReference,
} from "../../typings/Destination";
import "./DestinationConfigFields.scss";
import { EditIcon } from "../Icons";
import Tooltip from "../Tooltip/Tooltip";

const DestinationConfigFields = ({
  destination,
  type,
}: {
  destination?: Destination;
  type: DestinationTypeReference;
}) => {
  const [unlockedFields, setUnlockedFields] = useState<Record<string, boolean>>(
    {}
  );
  const [lastUnlockedField, setLastUnlockedField] = useState<string | null>(
    null
  );

  const inputRefs = useRef<Record<string, HTMLInputElement>>({});

  const unlockField = (key: string) => {
    setUnlockedFields((prev) => ({
      ...prev,
      [key]: !prev[key],
    }));
    setLastUnlockedField(key);

    if (inputRefs.current[key]) {
      inputRefs.current[key].value = "";
    }
  };

  useEffect(() => {
    if (lastUnlockedField && inputRefs.current[lastUnlockedField]) {
      inputRefs.current[lastUnlockedField].focus();
    }
  }, [lastUnlockedField]);

  return (
    <>
      {[...type.config_fields, ...type.credential_fields].map((field) => (
        <div key={field.key} className="destination-config-field">
          <label htmlFor={field.key}>
            {field.label}
            {field.required && <span className="required">*</span>}
          </label>
          {field.type === "text" && (
            <div className="input-container">
              <input
                ref={(el) => {
                  if (el) inputRefs.current[field.key] = el;
                }}
                type={
                  "sensitive" in field && field.sensitive ? "password" : "text"
                }
                placeholder={''}
                id={field.key}
                name={field.key}
                defaultValue={
                  "sensitive" in field && field.sensitive
                    ? unlockedFields[field.key]
                      ? ""
                      : destination?.credentials[field.key] || ""
                    : destination?.config[field.key] || ""
                }
                disabled={
                  "sensitive" in field && field.sensitive
                    ? destination?.credentials[field.key] &&
                      !unlockedFields[field.key]
                    : field.disabled
                }
                required={field.required}
                minLength={field.minlength}
                maxLength={field.maxlength}
                pattern={field.pattern}
              />
              {"sensitive" in field &&
              field.sensitive &&
              destination?.credentials[field.key] &&
              !unlockedFields[field.key] ? (
                <button type="button" onClick={() => unlockField(field.key)}>
                  <Tooltip content="Edit secret value" align="end">
                    <EditIcon />
                  </Tooltip>
                </button>
              ) : null}
            </div>
          )}
          {field.type === "checkbox" && (
            <input
              type="checkbox"
              id={field.key}
              name={field.key}
              defaultChecked={
                destination?.config[field.key] ??
                destination?.credentials[field.key] ??
                false
              }
              disabled={field.disabled}
              required={field.required}
            />
          )}
          {field.description && (
            <p className="description">{field.description}</p>
          )}
        </div>
      ))}
    </>
  );
};

export default DestinationConfigFields;

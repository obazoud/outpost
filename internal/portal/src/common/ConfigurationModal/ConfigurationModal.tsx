import { useEffect, useState } from "react";
import { createPortal } from "react-dom";
import { DestinationTypeReference } from "../../typings/Destination";
import { CollapseIcon } from "../Icons";
import Button from "../Button/Button";
import "./ConfigurationModal.scss";
import Markdown from "react-markdown";

const ConfigurationModal = ({
  type,
  onClose,
}: {
  type: DestinationTypeReference;
  onClose: () => void;
}) => {
  const [portalRef, setPortalRef] = useState<HTMLDivElement | null>(null);

  useEffect(() => {
    // Create portal container
    const portal = document.createElement("div");
    portal.id = "sidebar";
    document.body.appendChild(portal);
    document.body.classList.add("sidebar-open");
    setPortalRef(portal);
    return () => {
      // Cleanup on unmount
      portal.remove();
      document.body.classList.remove("sidebar-open");
    };
  }, []);

  if (!portalRef) {
    return null;
  }

  return createPortal(
    <>
      <Button minimal onClick={onClose} className="close-button">
        <CollapseIcon />
      </Button>
      <Markdown>{type.instructions}</Markdown>
    </>,
    portalRef
  );
};

export default ConfigurationModal;

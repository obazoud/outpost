import * as React from "react";
import * as ToastPrimitive from "@radix-ui/react-toast";
import "./Toast.scss";
import { CloseIcon, ErrorIcon, SuccessIcon } from "../Icons";

type ToastType = "success" | "error";

// Create an event emitter for toast messages
type ToastEvent = {
  type: ToastType;
  message: string;
};

const toastEventEmitter = {
  listeners: [] as ((event: ToastEvent) => void)[],
  emit(event: ToastEvent) {
    this.listeners.forEach((listener) => listener(event));
  },
  subscribe(listener: (event: ToastEvent) => void) {
    this.listeners.push(listener);
    return () => {
      this.listeners = this.listeners.filter((l) => l !== listener);
    };
  },
};

// The external function that can be called from anywhere
export const showToast = (type: ToastType, message: string) => {
  toastEventEmitter.emit({ type, message });
};

export const ToastProvider = ({ children }: { children: React.ReactNode }) => {
  const [open, setOpen] = React.useState(false);
  const [message, setMessage] = React.useState("");
  const [type, setType] = React.useState<ToastType>("success");

  React.useEffect(() => {
    const unsubscribe = toastEventEmitter.subscribe((event) => {
      setOpen(true);
      setType(event.type);
      setMessage(event.message);
      setOpen(true);
    });

    return () => unsubscribe();
  }, []);

  return (
    <>
      {children}
      <ToastPrimitive.Provider swipeDirection="right" duration={3000}>
        <ToastPrimitive.Root
          className={`ToastRoot ToastRoot__${type}`}
          open={open}
          onOpenChange={setOpen}
        >
          <ToastPrimitive.Title className="ToastTitle">
            <div className="ToastIcon">
              {type === "success" ? <SuccessIcon /> : <ErrorIcon />}
            </div>
            {message}
          </ToastPrimitive.Title>
          <ToastPrimitive.Close className="ToastClose" aria-label="Close">
            <CloseIcon />
          </ToastPrimitive.Close>
        </ToastPrimitive.Root>
        <ToastPrimitive.Viewport className="ToastViewport" />
      </ToastPrimitive.Provider>
    </>
  );
};

import { useEffect } from "react";
import Sentry from "./../../Sentry";

import "./ErrorBoundary.scss";
import { useLocation } from "react-router-dom";

export default ({ children }: { children: React.ReactNode }) => (
  <Sentry.ErrorBoundary
    fallback={({ resetError }) => {
      const location = useLocation();

      useEffect(() => {
        resetError();
      }, [location.pathname]);

      return (
        <div className="error-boundary">
          <h1 className="title-l">Something went wrong.</h1>
          <p className="body-s muted">Please try refreshing the page.</p>
        </div>
      );
    }}
  >
    {children}
  </Sentry.ErrorBoundary>
);

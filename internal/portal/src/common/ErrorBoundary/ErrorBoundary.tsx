import Sentry from "./../../Sentry";

import "./ErrorBoundary.scss";

export default ({ children }: { children: React.ReactNode }) => (
  <Sentry.ErrorBoundary
    fallback={
      <div className="error-boundary">
        <h1 className="title-l">Something went wrong.</h1>
        <p className="body-s muted">Please try refreshing the page.</p>
      </div>
    }
  >
    {children}
  </Sentry.ErrorBoundary>
);

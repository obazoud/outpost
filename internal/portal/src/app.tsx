import { useEffect, useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import DestinationList from "./scenes/DestinationsList/DestinationList";
import { SWRConfig } from "swr";

import "./global.scss";

export function App() {
  const token = useToken();
  const tenant = useTenant(token ?? undefined);

  if (!tenant) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {LOGO ? LOGO.indexOf("http") === 0 ? (
        <img src={LOGO} alt={ORGANIZATION_NAME} />
      ) : (
        <div dangerouslySetInnerHTML={{ __html: LOGO }} />
      ): null}
      <h1>{ORGANIZATION_NAME}</h1>

      <SWRConfig
        value={{
          fetcher: (path: string) =>
            fetch(`http://localhost:3333/api/v1/${tenant.id}/${path}`, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            }).then((res) => res.json()),
        }}
      >
        <BrowserRouter>
          <Routes>
            <Route path="/" Component={DestinationList} />
            <Route path="/new" element={<div>New Destination</div>} />
            <Route
              path="/:destination_id"
              element={<div>Specific Destination</div>}
            />
          </Routes>
        </BrowserRouter>
      </SWRConfig>
    </>
  );
}

function useToken() {
  const [token, setToken] = useState(sessionStorage.getItem("token"));

  useEffect(() => {
    const searchParams = new URLSearchParams(window.location.search);
    const token = searchParams.get("token");
    if (token) {
      setToken(token);
      sessionStorage.setItem("token", token);
      window.location.replace("/");
    }
  }, []);

  return token;
}

type Tenant = {
  id: string;
  created_at: string;
};

function useTenant(token?: string): Tenant | undefined {
  const [tenant, setTenant] = useState<Tenant | undefined>();

  useEffect(() => {
    run();

    async function run() {
      if (!token) {
        return;
      }
      const value = decodeJWT(token);
      if (!value.sub) {
        console.error("invalid token");
        return;
      }
      const tenantId = value.sub;
      const response = await fetch(`/api/v1/${tenantId}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        console.error(response.statusText);
        return;
      }
      const tenant = await response.json();
      setTenant(tenant);
    }
  }, [token]);

  return tenant;
}

function decodeJWT(token: string) {
  try {
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map(function (c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );
    return JSON.parse(jsonPayload);
  } catch (e) {
    console.error(e);
    return {};
  }
}

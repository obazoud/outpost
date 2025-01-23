import type { ZudokuConfig } from "zudoku";

const config: ZudokuConfig = {
  redirects: [
    {from: "/", to: "/docs"},
    {from: "/docs/guides", to: "/docs/guides/deployment"},
    {from: "/docs/references", to: "/docs/references/api"},
  ],
  page: {
    pageTitle: "",
    logoUrl: "/",
    logo: {
      src: {
        light: "/logo/outpost-logo-black.svg",
        dark: "/logo/outpost-logo-white.svg"
      },
      width: "110px",
    }
  },
  topNavigation: [
    { id: "docs", label: "Documentation", default: "docs" },
    // { id: "docs/api", label: "API Reference" },
    // { id: "/", label: "Website" },
  ],
  sidebar: {
    docs: [
      {
        type: "doc",
        label: "Overview",
        id: "docs",
      },
      {
        type: "doc",
        label: "Quickstart",
        id: "docs/quickstart",
      },
      {
        type: "doc",
        label: "Concepts",
        id: "docs/concepts",
      },
      {
        type: "doc",
        label: "Features",
        id: "docs/features",
      },
      // {
      //   type: "doc",
      //   label: "Configuration",
      //   id: "docs/configuration",
      // },
      {
        type: "category",
        label: "Guides",
        collapsed: false,
        collapsible: false,
        items: [
          {
            type: "doc",
            label: "Deployment",
            id: "docs/guides/deployment",
          },
          {
            type: "doc",
            label: "Dashboard Design",
            id: "docs/guides/dashboard-design",
          },
        ],
      },
      {
        type: "category",
        label: "References",
        collapsed: false,
        collapsible: false,
        items: [
          {
            type: "doc",
            label: "Configuration",
            id: "docs/references/configuration",
          },
          {
            type: "doc",
            label: "API",
            id: "docs/references/api",
          },
        ],
      },
    ],
  },

  // redirects: [{ from: "/", to: "/docs/introduction" }],
  apis: {
    type: "file",
    input: "./apis/openapi.yaml",
    navigationId: "docs/api",
  },
  docs: {
    files: "/pages/**/*.{md,mdx}",
  },
};

export default config;

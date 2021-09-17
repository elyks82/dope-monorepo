import { extendTheme } from "@chakra-ui/react"

export default extendTheme({
  colors: {
    transparent: "transparent",
    black: "#000",
    white: "#fff",
    gray: {
      50: "#f7fafc",
      900: "#171923",
    },
  },
  components: {
    Button: {
      variants: {
        solid: {
          background: "#EDEFEE",
          border: "2px solid #000000",
          borderRadius: "4px",
          boxShadow: "inset -1px -1px 0px rgba(0, 0, 0, 0.25), inset 1px 1px 0px rgba(255, 255, 255, 0.25)",
          height: "28px",
          fontSize: "12px",
          fontWeight: "400",
        }
      }
    },
    Table: {
      variants: {
        dope: {
          thead: {
            borderBottom: "2px solid #000",
          },
          th: {
            height: "34px",
            background: "#DEDEDD",
            textAlign: "center",
            verticalAlign: "middle",
            padding: 0,
            boxShadow: "inset -3px -3px 0px rgba(0, 0, 0, 0.25), inset 3px 3px 0px rgba(255, 255, 255, 0.25)",
            ':not(:last-child)': {
              borderRight: "2px solid #000",
            }
          },
          td: {
            padding: 0,
            verticalAlign: "middle",
            fontSize: "13px",
            textAlign: "center",
            height: "40px",
          },
          tr: {
            background: "#EDEFEE",
            textAlign: "center",
            _odd: {
              background: "#fff"
            }
          }
        }
      }
    }
  }
})
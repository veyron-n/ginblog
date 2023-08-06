export function isNull(v) {
  let result = typeof v === "undefined" || v === null || v === ''
  if (typeof v === "string") {
    result = v.trim() === ''
  }
  if (Array.isArray(v)) {
    result = v.length === 0
  }
  return result
}
# language_wizard
Language Wizard is a tiny, thread-safe i18n helper for Go. Store the active ISO language and a dictionary of strings, read with safe defaults, and hot-swap languages atomically. Wait for changes via a simple event channel, plug in custom logging for missing keys, and close cleanly.  Includes defensive copies and clear error types.

package imgui

// Flags for ImGui::InputText()
const (
	InputTextFlagsNone                = 0
	InputTextFlagsCharsDecimal        = 1 << 0  // Allow 0123456789.+-*/
	InputTextFlagsCharsHexadecimal    = 1 << 1  // Allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsUppercase      = 1 << 2  // Turn a..z into A..Z
	InputTextFlagsCharsNoBlank        = 1 << 3  // Filter out spaces, tabs
	InputTextFlagsAutoSelectAll       = 1 << 4  // Select entire text when first taking mouse focus
	InputTextFlagsEnterReturnsTrue    = 1 << 5  // Return 'true' when Enter is pressed (as opposed to when the value was modified)
	InputTextFlagsCallbackCompletion  = 1 << 6  // Callback on pressing TAB (for completion handling)
	InputTextFlagsCallbackHistory     = 1 << 7  // Callback on pressing Up/Down arrows (for history handling)
	InputTextFlagsCallbackAlways      = 1 << 8  // Callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackCharFilter  = 1 << 9  // Callback on character inputs to replace or discard them. Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsAllowTabInput       = 1 << 10 // Pressing TAB input a '\t' character into the text field
	InputTextFlagsCtrlEnterForNewLine = 1 << 11 // In multi-line mode, unfocus with Enter, add new line with Ctrl+Enter (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsNoHorizontalScroll  = 1 << 12 // Disable following the cursor horizontally
	InputTextFlagsAlwaysInsertMode    = 1 << 13 // Insert mode
	InputTextFlagsReadOnly            = 1 << 14 // Read-only mode
	InputTextFlagsPassword            = 1 << 15 // Password mode, display all characters as '*'
	InputTextFlagsNoUndoRedo          = 1 << 16 // Disable undo/redo. Note that input text owns the text data while active, if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsCharsScientific     = 1 << 17 // Allow 0123456789.+-*/eE (Scientific notation input)
	InputTextFlagsCallbackResize      = 1 << 18 // Callback on buffer capacity changes request (beyond 'buf_size' parameter value), allowing the string to grow. Notify when the string wants to be resized (for string types which hold a cache of their Size). You will be provided a new BufSize in the callback and NEED to honor it. (see misc/cpp/imgui_stdlib.h for an example of using this)
	// [Internal]
	InputTextFlagsMultiline = 1 << 20 // For internal use by InputTextMultiline()
)

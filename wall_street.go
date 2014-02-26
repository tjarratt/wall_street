package wall_street

// enums.go
type UndoEnum int
const (
	undoDelete UndoEnum = iota
	undoInsert UndoEnum = iota
	undoBegin UndoEnum = iota
	undoEnd UndoEnum = iota
)

type UndoList struct {
	next *UndoList
	start int // where the change took place
	end int
	text string // text to insert, if undoing a delete
	what UndoEnum
}

// the current undo list for the RL_LINE_BUFFER
var rl_undo_list *UndoList

// data structure for mapping textual names to code addresses
type FunMap struct {
	name string
	rlCommandFunc func() // rl_command_func_t
}

// /* **************************************************************** */
// /*								    */
// /*	     Functions available to bind to key sequences	    */
// /*								    */
// /* **************************************************************** */
// /* Bindable commands for numeric arguments. */
// extern int rl_digit_argument PARAMS((int, int));
// extern int rl_universal_argument PARAMS((int, int));

// /* Bindable commands for moving the cursor. */
// extern int rl_forward_byte PARAMS((int, int));
// extern int rl_forward_char PARAMS((int, int));
// extern int rl_forward PARAMS((int, int));
// extern int rl_backward_byte PARAMS((int, int));
// extern int rl_backward_char PARAMS((int, int));
// extern int rl_backward PARAMS((int, int));
// extern int rl_beg_of_line PARAMS((int, int));
// extern int rl_end_of_line PARAMS((int, int));
// extern int rl_forward_word PARAMS((int, int));
// extern int rl_backward_word PARAMS((int, int));
// extern int rl_refresh_line PARAMS((int, int));
// extern int rl_clear_screen PARAMS((int, int));
// extern int rl_skip_csi_sequence PARAMS((int, int));
// extern int rl_arrow_keys PARAMS((int, int));

// /* Bindable commands for inserting and deleting text. */
// extern int rl_insert PARAMS((int, int));
// extern int rl_quoted_insert PARAMS((int, int));
// extern int rl_tab_insert PARAMS((int, int));
// extern int rl_newline PARAMS((int, int));
// extern int rl_do_lowercase_version PARAMS((int, int));
// extern int rl_rubout PARAMS((int, int));
// extern int rl_delete PARAMS((int, int));
// extern int rl_rubout_or_delete PARAMS((int, int));
// extern int rl_delete_horizontal_space PARAMS((int, int));
// extern int rl_delete_or_show_completions PARAMS((int, int));
// extern int rl_insert_comment PARAMS((int, int));

// /* Bindable commands for changing case. */
// extern int rl_upcase_word PARAMS((int, int));
// extern int rl_downcase_word PARAMS((int, int));
// extern int rl_capitalize_word PARAMS((int, int));

// /* Bindable commands for transposing characters and words. */
// extern int rl_transpose_words PARAMS((int, int));
// extern int rl_transpose_chars PARAMS((int, int));

// /* Bindable commands for searching within a line. */
// extern int rl_char_search PARAMS((int, int));
// extern int rl_backward_char_search PARAMS((int, int));

// /* Bindable commands for readline's interface to the command history. */
// extern int rl_beginning_of_history PARAMS((int, int));
// extern int rl_end_of_history PARAMS((int, int));
// extern int rl_get_next_history PARAMS((int, int));
// extern int rl_get_previous_history PARAMS((int, int));

// /* Bindable commands for managing the mark and region. */
// extern int rl_set_mark PARAMS((int, int));
// extern int rl_exchange_point_and_mark PARAMS((int, int));

// /* Bindable commands to set the editing mode (emacs or vi). */
// extern int rl_vi_editing_mode PARAMS((int, int));
// extern int rl_emacs_editing_mode PARAMS((int, int));

// /* Bindable commands to change the insert mode (insert or overwrite) */
// extern int rl_overwrite_mode PARAMS((int, int));

// /* Bindable commands for managing key bindings. */
// extern int rl_re_read_init_file PARAMS((int, int));
// extern int rl_dump_functions PARAMS((int, int));
// extern int rl_dump_macros PARAMS((int, int));
// extern int rl_dump_variables PARAMS((int, int));

// /* Bindable commands for word completion. */
// extern int rl_complete PARAMS((int, int));
// extern int rl_possible_completions PARAMS((int, int));
// extern int rl_insert_completions PARAMS((int, int));
// extern int rl_old_menu_complete PARAMS((int, int));
// extern int rl_menu_complete PARAMS((int, int));
// extern int rl_backward_menu_complete PARAMS((int, int));

// /* Bindable commands for killing and yanking text, and managing the kill ring. */
// extern int rl_kill_word PARAMS((int, int));
// extern int rl_backward_kill_word PARAMS((int, int));
// extern int rl_kill_line PARAMS((int, int));
// extern int rl_backward_kill_line PARAMS((int, int));
// extern int rl_kill_full_line PARAMS((int, int));
// extern int rl_unix_word_rubout PARAMS((int, int));
// extern int rl_unix_filename_rubout PARAMS((int, int));
// extern int rl_unix_line_discard PARAMS((int, int));
// extern int rl_copy_region_to_kill PARAMS((int, int));
// extern int rl_kill_region PARAMS((int, int));
// extern int rl_copy_forward_word PARAMS((int, int));
// extern int rl_copy_backward_word PARAMS((int, int));
// extern int rl_yank PARAMS((int, int));
// extern int rl_yank_pop PARAMS((int, int));
// extern int rl_yank_nth_arg PARAMS((int, int));
// extern int rl_yank_last_arg PARAMS((int, int));
// /* Not available unless __CYGWIN__ is defined. */
// #ifdef __CYGWIN__
// extern int rl_paste_from_clipboard PARAMS((int, int));
// #endif

// /* Bindable commands for incremental searching. */
// extern int rl_reverse_search_history PARAMS((int, int));
// extern int rl_forward_search_history PARAMS((int, int));

// /* Bindable keyboard macro commands. */
// extern int rl_start_kbd_macro PARAMS((int, int));
// extern int rl_end_kbd_macro PARAMS((int, int));
// extern int rl_call_last_kbd_macro PARAMS((int, int));

// /* Bindable undo commands. */
// extern int rl_revert_line PARAMS((int, int));
// extern int rl_undo_command PARAMS((int, int));

// /* Bindable tilde expansion commands. */
// extern int rl_tilde_expand PARAMS((int, int));

// /* Bindable terminal control commands. */
// extern int rl_restart_output PARAMS((int, int));
// extern int rl_stop_output PARAMS((int, int));

// /* Miscellaneous bindable commands. */
// extern int rl_abort PARAMS((int, int));
// extern int rl_tty_status PARAMS((int, int));

// /* Bindable commands for incremental and non-incremental history searching. */
// extern int rl_history_search_forward PARAMS((int, int));
// extern int rl_history_search_backward PARAMS((int, int));
// extern int rl_noninc_forward_search PARAMS((int, int));
// extern int rl_noninc_reverse_search PARAMS((int, int));
// extern int rl_noninc_forward_search_again PARAMS((int, int));
// extern int rl_noninc_reverse_search_again PARAMS((int, int));

// /* Bindable command used when inserting a matching close character. */
// extern int rl_insert_close PARAMS((int, int));

// /* Not available unless READLINE_CALLBACKS is defined. */
// extern void rl_callback_handler_install PARAMS((const char *, rl_vcpfunc_t *));
// extern void rl_callback_read_char PARAMS((void));
// extern void rl_callback_handler_remove PARAMS((void));



// /* **************************************************************** */
// /*								    */
// /*			Well Published Functions		    */
// /*								    */
// /* **************************************************************** */

// /* Readline functions. */
// /* Read a line of input.  Prompt with PROMPT.  A NULL PROMPT means none. */
// extern char *readline PARAMS((const char *));

// extern int rl_set_prompt PARAMS((const char *));
// extern int rl_expand_prompt PARAMS((char *));

// extern int rl_initialize PARAMS((void));



// /* Utility functions to bind keys to readline commands. */
// extern int rl_add_defun PARAMS((const char *, rl_command_func_t *, int));
// extern int rl_bind_key PARAMS((int, rl_command_func_t *));
// extern int rl_bind_key_in_map PARAMS((int, rl_command_func_t *, Keymap));
// extern int rl_unbind_key PARAMS((int));
// extern int rl_unbind_key_in_map PARAMS((int, Keymap));
// extern int rl_bind_key_if_unbound PARAMS((int, rl_command_func_t *));
// extern int rl_bind_key_if_unbound_in_map PARAMS((int, rl_command_func_t *, Keymap));
// extern int rl_unbind_function_in_map PARAMS((rl_command_func_t *, Keymap));
// extern int rl_unbind_command_in_map PARAMS((const char *, Keymap));
// extern int rl_bind_keyseq PARAMS((const char *, rl_command_func_t *));
// extern int rl_bind_keyseq_in_map PARAMS((const char *, rl_command_func_t *, Keymap));
// extern int rl_bind_keyseq_if_unbound PARAMS((const char *, rl_command_func_t *));
// extern int rl_bind_keyseq_if_unbound_in_map PARAMS((const char *, rl_command_func_t *, Keymap));
// extern int rl_generic_bind PARAMS((int, const char *, char *, Keymap));

// extern char *rl_variable_value PARAMS((const char *));
// extern int rl_variable_bind PARAMS((const char *, const char *));



// /* Functions for manipulating keymaps. */
// extern Keymap rl_make_bare_keymap PARAMS((void));
// extern Keymap rl_copy_keymap PARAMS((Keymap));
// extern Keymap rl_make_keymap PARAMS((void));
// extern void rl_discard_keymap PARAMS((Keymap));

// extern Keymap rl_get_keymap_by_name PARAMS((const char *));
// extern char *rl_get_keymap_name PARAMS((Keymap));
// extern void rl_set_keymap PARAMS((Keymap));
// extern Keymap rl_get_keymap PARAMS((void));


// /* Functions for manipulating the funmap, which maps command names to functions. */
// extern int rl_add_funmap_entry PARAMS((const char *, rl_command_func_t *));
// extern const char **rl_funmap_names PARAMS((void));
// /* Undocumented, only used internally -- there is only one funmap, and this
//    function may be called only once. */
// extern void rl_initialize_funmap PARAMS((void));



// /* Utility functions for managing keyboard macros. */
// extern void rl_push_macro_input PARAMS((char *));



// /* Functions for undoing, from undo.c */
// extern void rl_add_undo PARAMS((enum undo_code, int, int, char *));
// extern void rl_free_undo_list PARAMS((void));
// extern int rl_do_undo PARAMS((void));
// extern int rl_begin_undo_group PARAMS((void));
// extern int rl_end_undo_group PARAMS((void));
// extern int rl_modifying PARAMS((int, int));



// /* Functions for redisplay. */
// extern void rl_redisplay PARAMS((void));
// extern int rl_on_new_line PARAMS((void));
// extern int rl_on_new_line_with_prompt PARAMS((void));
// extern int rl_forced_update_display PARAMS((void));
// extern int rl_clear_message PARAMS((void));
// extern int rl_reset_line_state PARAMS((void));
// extern int rl_crlf PARAMS((void));

// // assume we want to use varargs and prefer stdarg
// extern int rl_message (const char *, ...)  __attribute__((__format__ (printf, 1, 2)));

// // rl_show_char ???
// extern int rl_show_char PARAMS((int));

// /* Save and restore internal prompt redisplay information. */
// extern void rl_save_prompt PARAMS((void));
// extern void rl_restore_prompt PARAMS((void));

// /* Modifying text. */
// extern void rl_replace_line PARAMS((const char *, int));
// extern int rl_insert_text PARAMS((const char *));
// extern int rl_delete_text PARAMS((int, int));
// extern int rl_kill_text PARAMS((int, int));
// extern char *rl_copy_text PARAMS((int, int));


// /* Readline signal handling, from signals.c */
// extern int rl_set_signals PARAMS((void));
// extern int rl_clear_signals PARAMS((void));
// extern void rl_cleanup_after_signal PARAMS((void));
// extern void rl_reset_after_signal PARAMS((void));
// extern void rl_free_line_state PARAMS((void));
// extern void rl_echo_signal_char PARAMS((int));
// extern int rl_set_paren_blink_timeout PARAMS((int));

// /* Completion functions. */
// extern int rl_complete_internal PARAMS((int));
// extern void rl_display_match_list PARAMS((char **, int, int));
// extern char **rl_completion_matches PARAMS((const char *, rl_compentry_func_t *));
// extern char *rl_username_completion_function PARAMS((const char *, int));
// extern char *rl_filename_completion_function PARAMS((const char *, int));
// extern int rl_completion_mode PARAMS((rl_command_func_t *));

// +build go1.5,cgo
package plugin // import "collectd.org/plugin"

// #cgo CPPFLAGS: -DHAVE_CONFIG_H
// #cgo LDFLAGS: -ldl
// #include <stdlib.h>
// #include <dlfcn.h>
// #include "plugin.h"
//
// int (*register_read_ptr) (char const *group, char const *name,
//     plugin_read_cb callback,
//     cdtime_t interval,
//     user_data_t *ud) = NULL;
// int register_read_wrapper (char const *group, char const *name,
//     plugin_read_cb callback,
//     cdtime_t interval,
//     user_data_t *ud) {
//   if (register_read_ptr == NULL) {
//     void *hnd = dlopen(NULL, RTLD_LAZY);
//     register_read_ptr = dlsym(hnd, "plugin_register_complex_read");
//     dlclose(hnd);
//   }
//   return (*register_read_ptr) (group, name, callback, interval, ud);
// }
//
// int (*dispatch_values_ptr) (value_list_t const *vl);
// int dispatch_values_wrapper (value_list_t const *vl) {
//   if (dispatch_values_ptr == NULL) {
//     void *hnd = dlopen(NULL, RTLD_LAZY);
//     dispatch_values_ptr = dlsym(hnd, "plugin_dispatch_values");
//     dlclose(hnd);
//   }
//   return (*dispatch_values_ptr) (vl);
// }
//
// void value_list_add (value_list_t *vl, value_t v) {
//   value_t *tmp;
//   tmp = realloc (vl->values, (vl->values_len + 1));
//   if (tmp == NULL) {
//     errno = ENOMEM;
//     return;
//   }
//   vl->values = tmp;
//   vl->values[vl->values_len] = v;
//   vl->values_len++;
// }
//
// void value_list_add_gauge (value_list_t *vl, gauge_t g) {
//   value_t v = {.gauge = g};
//   value_list_add (vl, v);
// }
//
// void value_list_add_derive (value_list_t *vl, derive_t d) {
//   value_t v = {.derive = d};
//   value_list_add (vl, v);
// }
import "C"

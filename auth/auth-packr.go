// Code generated by github.com/gobuffalo/packr. DO NOT EDIT

package auth

import "github.com/gobuffalo/packr"

// You can use the "packr clean" command to clean up this,
// and any other packr generated files.
func init() {
	packr.PackJSONBytes("../auth/templates", "actions/auth.go.tmpl", "\"cGFja2FnZSBhY3Rpb25zCgppbXBvcnQgKAoJImRhdGFiYXNlL3NxbCIKCSJzdHJpbmdzIgoKICAie3sucGFja2FnZVBhdGh9fS9tb2RlbHMiCgkiZ2l0aHViLmNvbS9nb2J1ZmZhbG8vYnVmZmFsbyIKCSJnaXRodWIuY29tL2dvYnVmZmFsby9wb3AiCgkiZ2l0aHViLmNvbS9nb2J1ZmZhbG8vdmFsaWRhdGUiCgkiZ2l0aHViLmNvbS9wa2cvZXJyb3JzIgoJImdvbGFuZy5vcmcveC9jcnlwdG8vYmNyeXB0IgopCgovLyBBdXRoTmV3IGxvYWRzIHRoZSBzaWduaW4gcGFnZQpmdW5jIEF1dGhOZXcoYyBidWZmYWxvLkNvbnRleHQpIGVycm9yIHsKCWMuU2V0KCJ1c2VyIiwgbW9kZWxzLlVzZXJ7fSkKCXJldHVybiBjLlJlbmRlcigyMDAsIHIuSFRNTCgiYXV0aC9uZXcuaHRtbCIpKQp9CgovLyBBdXRoQ3JlYXRlIGF0dGVtcHRzIHRvIGxvZyB0aGUgdXNlciBpbiB3aXRoIGFuIGV4aXN0aW5nIGFjY291bnQuCmZ1bmMgQXV0aENyZWF0ZShjIGJ1ZmZhbG8uQ29udGV4dCkgZXJyb3IgewoJdSA6PSAmbW9kZWxzLlVzZXJ7fQoJaWYgZXJyIDo9IGMuQmluZCh1KTsgZXJyICE9IG5pbCB7CgkJcmV0dXJuIGVycm9ycy5XaXRoU3RhY2soZXJyKQoJfQoKCXR4IDo9IGMuVmFsdWUoInR4IikuKCpwb3AuQ29ubmVjdGlvbikKCgkvLyBmaW5kIGEgdXNlciB3aXRoIHRoZSBlbWFpbAoJZXJyIDo9IHR4LldoZXJlKCJlbWFpbCA9ID8iLCBzdHJpbmdzLlRvTG93ZXIodS5FbWFpbCkpLkZpcnN0KHUpCgoJLy8gaGVscGVyIGZ1bmN0aW9uIHRvIGhhbmRsZSBiYWQgYXR0ZW1wdHMKCWJhZCA6PSBmdW5jKCkgZXJyb3IgewoJCWMuU2V0KCJ1c2VyIiwgdSkKCQl2ZXJycyA6PSB2YWxpZGF0ZS5OZXdFcnJvcnMoKQoJCXZlcnJzLkFkZCgiZW1haWwiLCAiaW52YWxpZCBlbWFpbC9wYXNzd29yZCIpCgkJYy5TZXQoImVycm9ycyIsIHZlcnJzKQoJCXJldHVybiBjLlJlbmRlcig0MjIsIHIuSFRNTCgiYXV0aC9uZXcuaHRtbCIpKQoJfQoKCWlmIGVyciAhPSBuaWwgewoJCWlmIGVycm9ycy5DYXVzZShlcnIpID09IHNxbC5FcnJOb1Jvd3MgewoJCQkvLyBjb3VsZG4ndCBmaW5kIGFuIHVzZXIgd2l0aCB0aGUgc3VwcGxpZWQgZW1haWwgYWRkcmVzcy4KCQkJcmV0dXJuIGJhZCgpCgkJfQoJCXJldHVybiBlcnJvcnMuV2l0aFN0YWNrKGVycikKCX0KCgkvLyBjb25maXJtIHRoYXQgdGhlIGdpdmVuIHBhc3N3b3JkIG1hdGNoZXMgdGhlIGhhc2hlZCBwYXNzd29yZCBmcm9tIHRoZSBkYgoJZXJyID0gYmNyeXB0LkNvbXBhcmVIYXNoQW5kUGFzc3dvcmQoW11ieXRlKHUuUGFzc3dvcmRIYXNoKSwgW11ieXRlKHUuUGFzc3dvcmQpKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIGJhZCgpCgl9CgljLlNlc3Npb24oKS5TZXQoImN1cnJlbnRfdXNlcl9pZCIsIHUuSUQpCgljLkZsYXNoKCkuQWRkKCJzdWNjZXNzIiwgIldlbGNvbWUgQmFjayB0byBCdWZmYWxvISIpCgoJcmV0dXJuIGMuUmVkaXJlY3QoMzAyLCAiLyIpCn0KCi8vIEF1dGhEZXN0cm95IGNsZWFycyB0aGUgc2Vzc2lvbiBhbmQgbG9ncyBhIHVzZXIgb3V0CmZ1bmMgQXV0aERlc3Ryb3koYyBidWZmYWxvLkNvbnRleHQpIGVycm9yIHsKCWMuU2Vzc2lvbigpLkNsZWFyKCkKCWMuRmxhc2goKS5BZGQoInN1Y2Nlc3MiLCAiWW91IGhhdmUgYmVlbiBsb2dnZWQgb3V0ISIpCglyZXR1cm4gYy5SZWRpcmVjdCgzMDIsICIvIikKfQo=\"")
	packr.PackJSONBytes("../auth/templates", "actions/auth_test.go.tmpl", "\"cGFja2FnZSBhY3Rpb25zCgppbXBvcnQgKAogICJ7ey5wYWNrYWdlUGF0aH19L21vZGVscyIKKQoKZnVuYyAoYXMgKkFjdGlvblN1aXRlKSBUZXN0X0F1dGhfTmV3KCkgewoJcmVzIDo9IGFzLkhUTUwoIi9zaWduaW4iKS5HZXQoKQoJYXMuRXF1YWwoMjAwLCByZXMuQ29kZSkKCWFzLkNvbnRhaW5zKHJlcy5Cb2R5LlN0cmluZygpLCAiU2lnbiBJbiIpCn0KCmZ1bmMgKGFzICpBY3Rpb25TdWl0ZSkgVGVzdF9BdXRoX0NyZWF0ZSgpIHsKCXUgOj0gJm1vZGVscy5Vc2VyewoJCUVtYWlsOiAgICAgICAgICAgICAgICAibWFya0BleGFtcGxlLmNvbSIsCgkJUGFzc3dvcmQ6ICAgICAgICAgICAgICJwYXNzd29yZCIsCgkJUGFzc3dvcmRDb25maXJtYXRpb246ICJwYXNzd29yZCIsCgl9Cgl2ZXJycywgZXJyIDo9IHUuQ3JlYXRlKGFzLkRCKQoJYXMuTm9FcnJvcihlcnIpCglhcy5GYWxzZSh2ZXJycy5IYXNBbnkoKSkKCglyZXMgOj0gYXMuSFRNTCgiL3NpZ25pbiIpLlBvc3QodSkKCWFzLkVxdWFsKDMwMiwgcmVzLkNvZGUpCglhcy5FcXVhbCgiLyIsIHJlcy5Mb2NhdGlvbigpKQp9CgpmdW5jIChhcyAqQWN0aW9uU3VpdGUpIFRlc3RfQXV0aF9DcmVhdGVfVW5rbm93blVzZXIoKSB7Cgl1IDo9ICZtb2RlbHMuVXNlcnsKCQlFbWFpbDogICAgIm1hcmtAZXhhbXBsZS5jb20iLAoJCVBhc3N3b3JkOiAicGFzc3dvcmQiLAoJfQoJcmVzIDo9IGFzLkhUTUwoIi9zaWduaW4iKS5Qb3N0KHUpCglhcy5FcXVhbCg0MjIsIHJlcy5Db2RlKQoJYXMuQ29udGFpbnMocmVzLkJvZHkuU3RyaW5nKCksICJpbnZhbGlkIGVtYWlsL3Bhc3N3b3JkIikKfQoKZnVuYyAoYXMgKkFjdGlvblN1aXRlKSBUZXN0X0F1dGhfQ3JlYXRlX0JhZFBhc3N3b3JkKCkgewoJdSA6PSAmbW9kZWxzLlVzZXJ7CgkJRW1haWw6ICAgICAgICAgICAgICAgICJtYXJrQGV4YW1wbGUuY29tIiwKCQlQYXNzd29yZDogICAgICAgICAgICAgInBhc3N3b3JkIiwKCQlQYXNzd29yZENvbmZpcm1hdGlvbjogInBhc3N3b3JkIiwKCX0KCXZlcnJzLCBlcnIgOj0gdS5DcmVhdGUoYXMuREIpCglhcy5Ob0Vycm9yKGVycikKCWFzLkZhbHNlKHZlcnJzLkhhc0FueSgpKQoKCXUuUGFzc3dvcmQgPSAiYmFkIgoJcmVzIDo9IGFzLkhUTUwoIi9zaWduaW4iKS5Qb3N0KHUpCglhcy5FcXVhbCg0MjIsIHJlcy5Db2RlKQoJYXMuQ29udGFpbnMocmVzLkJvZHkuU3RyaW5nKCksICJpbnZhbGlkIGVtYWlsL3Bhc3N3b3JkIikKfQo=\"")
	packr.PackJSONBytes("../auth/templates", "actions/home_test.go.tmpl", "\"cGFja2FnZSBhY3Rpb25zCgppbXBvcnQgInt7LnBhY2thZ2VQYXRofX0vbW9kZWxzIgoKZnVuYyAoYXMgKkFjdGlvblN1aXRlKSBUZXN0X0hvbWVIYW5kbGVyKCkgewoJcmVzIDo9IGFzLkhUTUwoIi8iKS5HZXQoKQoJYXMuRXF1YWwoMjAwLCByZXMuQ29kZSkKCWFzLkNvbnRhaW5zKHJlcy5Cb2R5LlN0cmluZygpLCAiU2lnbiBJbiIpCn0KCmZ1bmMgKGFzICpBY3Rpb25TdWl0ZSkgVGVzdF9Ib21lSGFuZGxlcl9Mb2dnZWRJbigpIHsKCXUgOj0gJm1vZGVscy5Vc2VyewoJCUVtYWlsOiAgICAgICAgICAgICAgICAibWFya0BleGFtcGxlLmNvbSIsCgkJUGFzc3dvcmQ6ICAgICAgICAgICAgICJwYXNzd29yZCIsCgkJUGFzc3dvcmRDb25maXJtYXRpb246ICJwYXNzd29yZCIsCgl9Cgl2ZXJycywgZXJyIDo9IHUuQ3JlYXRlKGFzLkRCKQoJYXMuTm9FcnJvcihlcnIpCglhcy5GYWxzZSh2ZXJycy5IYXNBbnkoKSkKCWFzLlNlc3Npb24uU2V0KCJjdXJyZW50X3VzZXJfaWQiLCB1LklEKQoKCXJlcyA6PSBhcy5IVE1MKCIvIikuR2V0KCkKCWFzLkVxdWFsKDIwMCwgcmVzLkNvZGUpCglhcy5Db250YWlucyhyZXMuQm9keS5TdHJpbmcoKSwgIlNpZ24gT3V0IikKCglhcy5TZXNzaW9uLkNsZWFyKCkKCXJlcyA9IGFzLkhUTUwoIi8iKS5HZXQoKQoJYXMuRXF1YWwoMjAwLCByZXMuQ29kZSkKCWFzLkNvbnRhaW5zKHJlcy5Cb2R5LlN0cmluZygpLCAiU2lnbiBJbiIpCn0K\"")
	packr.PackJSONBytes("../auth/templates", "actions/users.go.tmpl", "\"cGFja2FnZSBhY3Rpb25zCgppbXBvcnQgKAogICJ7ey5wYWNrYWdlUGF0aH19L21vZGVscyIKCSJnaXRodWIuY29tL2dvYnVmZmFsby9idWZmYWxvIgoJImdpdGh1Yi5jb20vZ29idWZmYWxvL3BvcCIKCSJnaXRodWIuY29tL3BrZy9lcnJvcnMiCikKCmZ1bmMgVXNlcnNOZXcoYyBidWZmYWxvLkNvbnRleHQpIGVycm9yIHsKCXUgOj0gbW9kZWxzLlVzZXJ7fQoJYy5TZXQoInVzZXIiLCB1KQoJcmV0dXJuIGMuUmVuZGVyKDIwMCwgci5IVE1MKCJ1c2Vycy9uZXcuaHRtbCIpKQp9CgovLyBVc2Vyc0NyZWF0ZSByZWdpc3RlcnMgYSBuZXcgdXNlciB3aXRoIHRoZSBhcHBsaWNhdGlvbi4KZnVuYyBVc2Vyc0NyZWF0ZShjIGJ1ZmZhbG8uQ29udGV4dCkgZXJyb3IgewoJdSA6PSAmbW9kZWxzLlVzZXJ7fQoJaWYgZXJyIDo9IGMuQmluZCh1KTsgZXJyICE9IG5pbCB7CgkJcmV0dXJuIGVycm9ycy5XaXRoU3RhY2soZXJyKQoJfQoKCXR4IDo9IGMuVmFsdWUoInR4IikuKCpwb3AuQ29ubmVjdGlvbikKCXZlcnJzLCBlcnIgOj0gdS5DcmVhdGUodHgpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gZXJyb3JzLldpdGhTdGFjayhlcnIpCgl9CgoJaWYgdmVycnMuSGFzQW55KCkgewoJCWMuU2V0KCJ1c2VyIiwgdSkKCQljLlNldCgiZXJyb3JzIiwgdmVycnMpCgkJcmV0dXJuIGMuUmVuZGVyKDIwMCwgci5IVE1MKCJ1c2Vycy9uZXcuaHRtbCIpKQoJfQoKCWMuU2Vzc2lvbigpLlNldCgiY3VycmVudF91c2VyX2lkIiwgdS5JRCkKCWMuRmxhc2goKS5BZGQoInN1Y2Nlc3MiLCAiV2VsY29tZSB0byBCdWZmYWxvISIpCgoJcmV0dXJuIGMuUmVkaXJlY3QoMzAyLCAiLyIpCn0KCi8vIFNldEN1cnJlbnRVc2VyIGF0dGVtcHRzIHRvIGZpbmQgYSB1c2VyIGJhc2VkIG9uIHRoZSBjdXJyZW50X3VzZXJfaWQKLy8gaW4gdGhlIHNlc3Npb24uIElmIG9uZSBpcyBmb3VuZCBpdCBpcyBzZXQgb24gdGhlIGNvbnRleHQuCmZ1bmMgU2V0Q3VycmVudFVzZXIobmV4dCBidWZmYWxvLkhhbmRsZXIpIGJ1ZmZhbG8uSGFuZGxlciB7CglyZXR1cm4gZnVuYyhjIGJ1ZmZhbG8uQ29udGV4dCkgZXJyb3IgewoJCWlmIHVpZCA6PSBjLlNlc3Npb24oKS5HZXQoImN1cnJlbnRfdXNlcl9pZCIpOyB1aWQgIT0gbmlsIHsKCQkJdSA6PSAmbW9kZWxzLlVzZXJ7fQoJCQl0eCA6PSBjLlZhbHVlKCJ0eCIpLigqcG9wLkNvbm5lY3Rpb24pCgkJCWVyciA6PSB0eC5GaW5kKHUsIHVpZCkKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlyZXR1cm4gZXJyb3JzLldpdGhTdGFjayhlcnIpCgkJCX0KCQkJYy5TZXQoImN1cnJlbnRfdXNlciIsIHUpCgkJfQoJCXJldHVybiBuZXh0KGMpCgl9Cn0KCi8vIEF1dGhvcml6ZSByZXF1aXJlIGEgdXNlciBiZSBsb2dnZWQgaW4gYmVmb3JlIGFjY2Vzc2luZyBhIHJvdXRlCmZ1bmMgQXV0aG9yaXplKG5leHQgYnVmZmFsby5IYW5kbGVyKSBidWZmYWxvLkhhbmRsZXIgewoJcmV0dXJuIGZ1bmMoYyBidWZmYWxvLkNvbnRleHQpIGVycm9yIHsKCQlpZiB1aWQgOj0gYy5TZXNzaW9uKCkuR2V0KCJjdXJyZW50X3VzZXJfaWQiKTsgdWlkID09IG5pbCB7CgkJCWMuRmxhc2goKS5BZGQoImRhbmdlciIsICJZb3UgbXVzdCBiZSBhdXRob3JpemVkIHRvIHNlZSB0aGF0IHBhZ2UiKQoJCQlyZXR1cm4gYy5SZWRpcmVjdCgzMDIsICIvIikKCQl9CgkJcmV0dXJuIG5leHQoYykKCX0KfQo=\"")
	packr.PackJSONBytes("../auth/templates", "actions/users_test.go.tmpl", "\"cGFja2FnZSBhY3Rpb25zCgppbXBvcnQgKAogICJ7ey5wYWNrYWdlUGF0aH19L21vZGVscyIKKQoKZnVuYyAoYXMgKkFjdGlvblN1aXRlKSBUZXN0X1VzZXJzX05ldygpIHsKCXJlcyA6PSBhcy5IVE1MKCIvdXNlcnMvbmV3IikuR2V0KCkKCWFzLkVxdWFsKDIwMCwgcmVzLkNvZGUpCn0KCmZ1bmMgKGFzICpBY3Rpb25TdWl0ZSkgVGVzdF9Vc2Vyc19DcmVhdGUoKSB7Cgljb3VudCwgZXJyIDo9IGFzLkRCLkNvdW50KCJ1c2VycyIpCglhcy5Ob0Vycm9yKGVycikKCWFzLkVxdWFsKDAsIGNvdW50KQoKCXUgOj0gJm1vZGVscy5Vc2VyewoJCUVtYWlsOiAgICAgICAgICAgICAgICAibWFya0BleGFtcGxlLmNvbSIsCgkJUGFzc3dvcmQ6ICAgICAgICAgICAgICJwYXNzd29yZCIsCgkJUGFzc3dvcmRDb25maXJtYXRpb246ICJwYXNzd29yZCIsCgl9CgoJcmVzIDo9IGFzLkhUTUwoIi91c2VycyIpLlBvc3QodSkKCWFzLkVxdWFsKDMwMiwgcmVzLkNvZGUpCgoJY291bnQsIGVyciA9IGFzLkRCLkNvdW50KCJ1c2VycyIpCglhcy5Ob0Vycm9yKGVycikKCWFzLkVxdWFsKDEsIGNvdW50KQp9Cg==\"")
	packr.PackJSONBytes("../auth/templates", "models/user_test.go.tmpl", "\"cGFja2FnZSBtb2RlbHNfdGVzdAoKaW1wb3J0ICgKICAie3sucGFja2FnZVBhdGh9fS9tb2RlbHMiCikKCmZ1bmMgKG1zICpNb2RlbFN1aXRlKSBUZXN0X1VzZXJfQ3JlYXRlKCkgewoJY291bnQsIGVyciA6PSBtcy5EQi5Db3VudCgidXNlcnMiKQoJbXMuTm9FcnJvcihlcnIpCgltcy5FcXVhbCgwLCBjb3VudCkKCgl1IDo9ICZtb2RlbHMuVXNlcnsKCQlFbWFpbDogICAgICAgICAgICAgICAgIm1hcmtAZXhhbXBsZS5jb20iLAoJCVBhc3N3b3JkOiAgICAgICAgICAgICAicGFzc3dvcmQiLAoJCVBhc3N3b3JkQ29uZmlybWF0aW9uOiAicGFzc3dvcmQiLAoJfQoJbXMuWmVybyh1LlBhc3N3b3JkSGFzaCkKCgl2ZXJycywgZXJyIDo9IHUuQ3JlYXRlKG1zLkRCKQoJbXMuTm9FcnJvcihlcnIpCgltcy5GYWxzZSh2ZXJycy5IYXNBbnkoKSkKCW1zLk5vdFplcm8odS5QYXNzd29yZEhhc2gpCgoJY291bnQsIGVyciA9IG1zLkRCLkNvdW50KCJ1c2VycyIpCgltcy5Ob0Vycm9yKGVycikKCW1zLkVxdWFsKDEsIGNvdW50KQp9CgpmdW5jIChtcyAqTW9kZWxTdWl0ZSkgVGVzdF9Vc2VyX0NyZWF0ZV9WYWxpZGF0aW9uRXJyb3JzKCkgewoJY291bnQsIGVyciA6PSBtcy5EQi5Db3VudCgidXNlcnMiKQoJbXMuTm9FcnJvcihlcnIpCgltcy5FcXVhbCgwLCBjb3VudCkKCgl1IDo9ICZtb2RlbHMuVXNlcnsKCQlQYXNzd29yZDogInBhc3N3b3JkIiwKCX0KCW1zLlplcm8odS5QYXNzd29yZEhhc2gpCgoJdmVycnMsIGVyciA6PSB1LkNyZWF0ZShtcy5EQikKCW1zLk5vRXJyb3IoZXJyKQoJbXMuVHJ1ZSh2ZXJycy5IYXNBbnkoKSkKCgljb3VudCwgZXJyID0gbXMuREIuQ291bnQoInVzZXJzIikKCW1zLk5vRXJyb3IoZXJyKQoJbXMuRXF1YWwoMCwgY291bnQpCn0KCmZ1bmMgKG1zICpNb2RlbFN1aXRlKSBUZXN0X1VzZXJfQ3JlYXRlX1VzZXJFeGlzdHMoKSB7Cgljb3VudCwgZXJyIDo9IG1zLkRCLkNvdW50KCJ1c2VycyIpCgltcy5Ob0Vycm9yKGVycikKCW1zLkVxdWFsKDAsIGNvdW50KQoKCXUgOj0gJm1vZGVscy5Vc2VyewoJCUVtYWlsOiAgICAgICAgICAgICAgICAibWFya0BleGFtcGxlLmNvbSIsCgkJUGFzc3dvcmQ6ICAgICAgICAgICAgICJwYXNzd29yZCIsCgkJUGFzc3dvcmRDb25maXJtYXRpb246ICJwYXNzd29yZCIsCgl9Cgltcy5aZXJvKHUuUGFzc3dvcmRIYXNoKQoKCXZlcnJzLCBlcnIgOj0gdS5DcmVhdGUobXMuREIpCgltcy5Ob0Vycm9yKGVycikKCW1zLkZhbHNlKHZlcnJzLkhhc0FueSgpKQoJbXMuTm90WmVybyh1LlBhc3N3b3JkSGFzaCkKCgljb3VudCwgZXJyID0gbXMuREIuQ291bnQoInVzZXJzIikKCW1zLk5vRXJyb3IoZXJyKQoJbXMuRXF1YWwoMSwgY291bnQpCgoJdSA9ICZtb2RlbHMuVXNlcnsKCQlFbWFpbDogICAgIm1hcmtAZXhhbXBsZS5jb20iLAoJCVBhc3N3b3JkOiAicGFzc3dvcmQiLAoJfQoJdmVycnMsIGVyciA9IHUuQ3JlYXRlKG1zLkRCKQoJbXMuTm9FcnJvcihlcnIpCgltcy5UcnVlKHZlcnJzLkhhc0FueSgpKQoKCWNvdW50LCBlcnIgPSBtcy5EQi5Db3VudCgidXNlcnMiKQoJbXMuTm9FcnJvcihlcnIpCgltcy5FcXVhbCgxLCBjb3VudCkKfQo=\"")
	packr.PackJSONBytes("../auth/templates", "templates/auth/new.html.tmpl", "\"PHN0eWxlPgogIC5hdXRoLXdyYXBwZXJ7CiAgICBoZWlnaHQ6IDEwMCU7CiAgICBkaXNwbGF5OiBmbGV4OwogICAgYWxpZ24taXRlbXM6IGNlbnRlcjsKICAgIGp1c3RpZnktY29udGVudDogY2VudGVyOwogIH0KCiAgLmF1dGgtd3JhcHBlciAuc2lnbi1mb3JtewogICAgbWF4LXdpZHRoOiAzNTBweDsKICAgIHdpZHRoOiAxMDAlOwogICAgcGFkZGluZzogMCAyMHB4OwogIH0KCiAgLmF1dGgtd3JhcHBlciBoMXttYXJnaW4tYm90dG9tOiAyMHB4O30KPC9zdHlsZT4KCjxkaXYgY2xhc3M9ImF1dGgtd3JhcHBlciI+CiAgPGRpdiBjbGFzcz0ic2lnbi1mb3JtIj4KICAgIDxoMT5TaWduIEluPC9oMT4KCiAgICA8JT0gZm9ybV9mb3IodXNlciwge2FjdGlvbjogc2lnbmluUGF0aCgpfSkgeyAlPgogICAgICA8JT0gZi5JbnB1dFRhZygiRW1haWwiKSAlPgogICAgICA8JT0gZi5JbnB1dFRhZygiUGFzc3dvcmQiLCB7dHlwZTogInBhc3N3b3JkIn0pICU+CiAgICAgIDxidXR0b24gY2xhc3M9ImJ0biBidG4tc3VjY2VzcyI+U2lnbiBJbiE8L2J1dHRvbj4KICAgIDwlIH0gJT4KICA8L2Rpdj4KPC9kaXY+\"")
	packr.PackJSONBytes("../auth/templates", "templates/index.html.tmpl", "\"PHN0eWxlPgogIC5hdXRoLWNlbnRlcnsKICAgIGRpc3BsYXk6IGZsZXg7CiAgICBhbGlnbi1pdGVtczogY2VudGVyOwogICAganVzdGlmeS1jb250ZW50OiBjZW50ZXI7CiAgICBoZWlnaHQ6IDEwMCU7CiAgfQogIC5zaWduLWluLWJ0bnsKICAgIG1hcmdpbi1yaWdodDogMTBweDsKICB9Cjwvc3R5bGU+Cgo8ZGl2IGNsYXNzPSJhdXRoLWNlbnRlciI+CiAgPCU9IGlmIChjdXJyZW50X3VzZXIpIHsgJT4KICAgIDxoMT48JT0gY3VycmVudF91c2VyLkVtYWlsICU+PC9oMT4KICAgIDxhIGhyZWY9Ii9zaWdub3V0IiBkYXRhLW1ldGhvZD0iREVMRVRFIj5TaWduIE91dDwvYT4KICA8JSB9IGVsc2UgeyAlPgogICAgPGEgaHJlZj0iL3NpZ25pbiIgY2xhc3M9ImJ0biBidG4tcHJpbWFyeSI+U2lnbiBJbjwvYT4KICAgIDxhIGhyZWY9Ii91c2Vycy9uZXciIGNsYXNzPSJidG4gYnRuLXN1Y2Nlc3MiPlJlZ2lzdGVyPC9hPgogIDwlIH0gJT4KPC9kaXY+Cg==\"")
	packr.PackJSONBytes("../auth/templates", "templates/users/new.html.tmpl", "\"PHN0eWxlPgogIC5hdXRoLXdyYXBwZXJ7CiAgICBoZWlnaHQ6IDEwMCU7CiAgICBkaXNwbGF5OiBmbGV4OwogICAgYWxpZ24taXRlbXM6IGNlbnRlcjsKICAgIGp1c3RpZnktY29udGVudDogY2VudGVyOwogIH0KCiAgLmF1dGgtd3JhcHBlciAuc2lnbi1mb3JtewogICAgbWF4LXdpZHRoOiAzNTBweDsKICAgIHdpZHRoOiAxMDAlOwogICAgcGFkZGluZzogMCAyMHB4OwogIH0KCiAgLmF1dGgtd3JhcHBlciBoMXttYXJnaW4tYm90dG9tOiAyMHB4O30KPC9zdHlsZT4KICAKPGRpdiBjbGFzcz0iYXV0aC13cmFwcGVyIj4KICA8ZGl2IGNsYXNzPSJzaWduLWZvcm0iPgogICAgPGgxPlJlZ2lzdGVyPC9oMT4KCiAgICA8JT0gZm9ybV9mb3IodXNlciwge2FjdGlvbjogdXNlcnNQYXRoKCl9KSB7ICU+CiAgICAgIDwlPSBmLklucHV0VGFnKCJFbWFpbCIpICU+CiAgICAgIDwlPSBmLklucHV0VGFnKCJQYXNzd29yZCIsIHt0eXBlOiAicGFzc3dvcmQifSkgJT4KICAgICAgPCU9IGYuSW5wdXRUYWcoIlBhc3N3b3JkQ29uZmlybWF0aW9uIiwge3R5cGU6ICJwYXNzd29yZCJ9KSAlPgogICAgICAKICAgICAgPGJ1dHRvbiBjbGFzcz0iYnRuIGJ0bi1zdWNjZXNzIj5SZWdpc3RlciE8L2J1dHRvbj4KICAgIDwlIH0gJT4KICA8L2Rpdj4KPC9kaXY+\"")
}

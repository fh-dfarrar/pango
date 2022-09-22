/*
Package saml is the client.Device.SamlProfile namespace.

For Panorama, there are two possibilities:  managing this object on Panorama
itself or inside of a Template.

To manage objects save on Panorama, leave "tmpl", "ts", and "vsys" params empty.

To manage objects in a template, specify the template name and the vsys (if
unspecified, defaults to "shared").

PAN-OS 8.0+

Normalized object:  Entry
*/
package saml

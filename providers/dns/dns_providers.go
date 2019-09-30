package dns

import (
	"fmt"

	"github.com/arno01/lego/v3/challenge"
	"github.com/arno01/lego/v3/challenge/dns01"
	"github.com/arno01/lego/v3/providers/dns/acmedns"
	"github.com/arno01/lego/v3/providers/dns/alidns"
	"github.com/arno01/lego/v3/providers/dns/auroradns"
	"github.com/arno01/lego/v3/providers/dns/azure"
	"github.com/arno01/lego/v3/providers/dns/bindman"
	"github.com/arno01/lego/v3/providers/dns/bluecat"
	"github.com/arno01/lego/v3/providers/dns/cloudflare"
	"github.com/arno01/lego/v3/providers/dns/cloudns"
	"github.com/arno01/lego/v3/providers/dns/cloudxns"
	"github.com/arno01/lego/v3/providers/dns/conoha"
	"github.com/arno01/lego/v3/providers/dns/designate"
	"github.com/arno01/lego/v3/providers/dns/digitalocean"
	"github.com/arno01/lego/v3/providers/dns/dnsimple"
	"github.com/arno01/lego/v3/providers/dns/dnsmadeeasy"
	"github.com/arno01/lego/v3/providers/dns/dnspod"
	"github.com/arno01/lego/v3/providers/dns/dode"
	"github.com/arno01/lego/v3/providers/dns/dreamhost"
	"github.com/arno01/lego/v3/providers/dns/duckdns"
	"github.com/arno01/lego/v3/providers/dns/dyn"
	"github.com/arno01/lego/v3/providers/dns/easydns"
	"github.com/arno01/lego/v3/providers/dns/exec"
	"github.com/arno01/lego/v3/providers/dns/exoscale"
	"github.com/arno01/lego/v3/providers/dns/fastdns"
	"github.com/arno01/lego/v3/providers/dns/gandi"
	"github.com/arno01/lego/v3/providers/dns/gandiv5"
	"github.com/arno01/lego/v3/providers/dns/gcloud"
	"github.com/arno01/lego/v3/providers/dns/glesys"
	"github.com/arno01/lego/v3/providers/dns/godaddy"
	"github.com/arno01/lego/v3/providers/dns/hostingde"
	"github.com/arno01/lego/v3/providers/dns/httpreq"
	"github.com/arno01/lego/v3/providers/dns/iij"
	"github.com/arno01/lego/v3/providers/dns/inwx"
	"github.com/arno01/lego/v3/providers/dns/joker"
	"github.com/arno01/lego/v3/providers/dns/lightsail"
	"github.com/arno01/lego/v3/providers/dns/linode"
	"github.com/arno01/lego/v3/providers/dns/linodev4"
	"github.com/arno01/lego/v3/providers/dns/liquidweb"
	"github.com/arno01/lego/v3/providers/dns/mydnsjp"
	"github.com/arno01/lego/v3/providers/dns/namecheap"
	"github.com/arno01/lego/v3/providers/dns/namedotcom"
	"github.com/arno01/lego/v3/providers/dns/namesilo"
	"github.com/arno01/lego/v3/providers/dns/netcup"
	"github.com/arno01/lego/v3/providers/dns/nifcloud"
	"github.com/arno01/lego/v3/providers/dns/ns1"
	"github.com/arno01/lego/v3/providers/dns/oraclecloud"
	"github.com/arno01/lego/v3/providers/dns/otc"
	"github.com/arno01/lego/v3/providers/dns/ovh"
	"github.com/arno01/lego/v3/providers/dns/pdns"
	"github.com/arno01/lego/v3/providers/dns/rackspace"
	"github.com/arno01/lego/v3/providers/dns/rfc2136"
	"github.com/arno01/lego/v3/providers/dns/route53"
	"github.com/arno01/lego/v3/providers/dns/sakuracloud"
	"github.com/arno01/lego/v3/providers/dns/selectel"
	"github.com/arno01/lego/v3/providers/dns/stackpath"
	"github.com/arno01/lego/v3/providers/dns/transip"
	"github.com/arno01/lego/v3/providers/dns/vegadns"
	"github.com/arno01/lego/v3/providers/dns/versio"
	"github.com/arno01/lego/v3/providers/dns/vscale"
	"github.com/arno01/lego/v3/providers/dns/vultr"
	"github.com/arno01/lego/v3/providers/dns/zoneee"
)

// NewDNSChallengeProviderByName Factory for DNS providers
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProvider()
	case "alidns":
		return alidns.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "bindman":
		return bindman.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudns":
		return cloudns.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "conoha":
		return conoha.NewDNSProvider()
	case "designate":
		return designate.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnsimple":
		return dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		return dnspod.NewDNSProvider()
	case "dode":
		return dode.NewDNSProvider()
	case "dreamhost":
		return dreamhost.NewDNSProvider()
	case "duckdns":
		return duckdns.NewDNSProvider()
	case "dyn":
		return dyn.NewDNSProvider()
	case "fastdns":
		return fastdns.NewDNSProvider()
	case "easydns":
		return easydns.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "hostingde":
		return hostingde.NewDNSProvider()
	case "httpreq":
		return httpreq.NewDNSProvider()
	case "iij":
		return iij.NewDNSProvider()
	case "inwx":
		return inwx.NewDNSProvider()
	case "joker":
		return joker.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode":
		return linode.NewDNSProvider()
	case "linodev4":
		return linodev4.NewDNSProvider()
	case "liquidweb":
		return liquidweb.NewDNSProvider()
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "namesilo":
		return namesilo.NewDNSProvider()
	case "netcup":
		return netcup.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "oraclecloud":
		return oraclecloud.NewDNSProvider()
	case "otc":
		return otc.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "stackpath":
		return stackpath.NewDNSProvider()
	case "selectel":
		return selectel.NewDNSProvider()
	case "transip":
		return transip.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	case "versio":
		return versio.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "vscale":
		return vscale.NewDNSProvider()
	case "zoneee":
		return zoneee.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}

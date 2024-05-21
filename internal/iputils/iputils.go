package iputils

import (
	"errors"
	"net"
	"strings"
)

// Result struct to encapsulate the processed result, warnings, and errors
type Result struct {
	Result   []string
	Warnings []string
	Errors   []string
}

// ParseIpCidr parses the input string into IPs and CIDRs based on the delimiter.
func parseIpCidr(input string) (ips []string, cidrs []string, err error) {
	parts := strings.Split(input, ";")
	if len(parts) != 2 {
		return nil, nil, errors.New("invalid ipCidr format")
	}


	///this is a test
	
	ipPart := strings.TrimSpace(parts[0])
	cidrPart := strings.TrimSpace(parts[1])

	if ipPart != "" {
		ips = strings.Split(ipPart, ",")
		for i := range ips {
			ips[i] = strings.TrimSpace(ips[i])
		}
	}

	if cidrPart != "" {
		cidrs = strings.Split(cidrPart, ",")
		for i := range cidrs {
			cidrs[i] = strings.TrimSpace(cidrs[i])
		}
	}

	return ips, cidrs, nil
}

// ValidateIPs validates if the given IPs are valid.
func validateIPs(ips []string) bool {
	for _, ip := range ips {
		if net.ParseIP(ip) == nil {
			return false
		}
	}
	return true
}

// ValidateCIDRs validates if the given CIDRs are valid.
func validateCIDRs(cidrs []string) bool {
	for _, cidr := range cidrs {
		if _, _, err := net.ParseCIDR(cidr); err != nil {
			return false
		}
	}
	return true
}

// CheckCidrsOverlap checks if there is an overlap between any CIDRs in the two slices.
func checkCidrsOverlap(cidrBlocks1, cidrBlocks2 []string) bool {
	for _, cidr1 := range cidrBlocks1 {
		_, ipnet1, _ := net.ParseCIDR(cidr1)
		for _, cidr2 := range cidrBlocks2 {
			_, ipnet2, _ := net.ParseCIDR(cidr2)
			if ipnet1.Contains(ipnet2.IP) || ipnet2.Contains(ipnet1.IP) {
				return true
			}
		}
	}
	return false
}

// ProcessIpCidr processes the IP and CIDR inputs, returning warnings if any, or the combined valid list.
func ProcessIpCidr(sourceAddressPrefixes []string, ipCidr string) (result Result) {
	// Parse sourceAddressPrefixes into IPs and CIDRs
	var sourceIps []string
	var sourceCidrs []string
	for _, prefix := range sourceAddressPrefixes {
		if strings.Contains(prefix, "/") {
			sourceCidrs = append(sourceCidrs, prefix)
		} else {
			sourceIps = append(sourceIps, prefix)
		}
	}

	// Parse ipCidr into IPs and CIDRs
	ips, cidrs, parseErr := parseIpCidr(ipCidr)
	if parseErr != nil {
		result.Errors = append(result.Errors, parseErr.Error())
		return result
	}

	// Validate IPs
	if !validateIPs(ips) {
		result.Errors = append(result.Errors, "invalid IP addresses")
	}

	// Validate CIDRs
	if !validateCIDRs(cidrs) {
		result.Errors = append(result.Errors, "invalid CIDR blocks")
	}

	// Check for overlapping CIDRs
	if checkCidrsOverlap(sourceCidrs, cidrs) {
		result.Errors = append(result.Errors, "overlapping CIDR blocks detected")
	}

	if len(result.Errors) > 0 {
		return result
	}

	// Combine and deduplicate IPs and CIDRs
	allIps := append(sourceIps, ips...)
	allCidrs := append(sourceCidrs, cidrs...)

	allIps = removeDuplicates(allIps)
	allCidrs = removeDuplicates(allCidrs)

	// Check for duplicate IPs and add warnings
	for _, ip := range ips {
		if containsString(sourceIps, ip) {
			result.Warnings = append(result.Warnings, "duplicate IP address: "+ip)
		}
	}

	result.Result = append(allIps, allCidrs...)
	return result
}

// Remove duplicates from a slice of strings.
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			continue
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// ContainsString checks if a slice contains a given string
func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

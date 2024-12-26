package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

type AcclItem struct {
	Name                string     `json:"name"`
	Namespace           string     `json:"namespace"`
	Creator             string     `json:"creator"`
	Priority            int        `json:"priority"`
	GPUs                int        `json:"gpus"`
	Image               string     `json:"image"`
	Status              AcclStatus `json:"status,omitempty"`
	SchedulingGroupName string     `json:"schedulingGroupName"`
	ReconcileStatus     string     `json:"reconcileStatus,omitempty"`

	FlavorName string `json:"flavorName"`
	MafVersion string `json:"mafVersion"`
}

type AcclStatus struct {
	Phase          string `json:"phase,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
}

/* // TODO
func PrintAcclKubectlStyle(data io.ReadCloser) error {
	var acclItems interface{}
	if err := json.NewDecoder(data).Decode(&acclItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if acclItems == nil {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tPRIORITY\tGPUS\tSTATUS\tPHASE\tIMAGE")

	switch items := acclItems.(type) {
	case []interface{}:
		for _, item := range items {
			printAccl(writer, item.(map[string]interface{}))
		}
	case map[string]interface{}:
		//printAccl(writer, items)
		fmt.Println(acclItems.(AcclItem))
	default:
		fmt.Println("Unexpected JSON format.")
	}

	return nil
}
*/

func PrintAcclK8sStyle(data io.ReadCloser) error {
	var acclItem *AcclItem
	if err := json.NewDecoder(data).Decode(&acclItem); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if acclItem == nil {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tSCHEDULING_GROUP_NAME\tPRIORITY\tGPUS\tSTATUS\tPHASE\tIMAGE")

	printAccl(writer, *acclItem)

	return nil
}

func PrintAcclListK8sStyle(data io.ReadCloser) error {
	var acclItems []AcclItem
	if err := json.NewDecoder(data).Decode(&acclItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(acclItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tSCHEDULING_GROUP_NAME\tPRIORITY\tGPUS\tSTATUS\tPHASE\tIMAGE")

	for _, item := range acclItems {
		printAccl(writer, item)
	}

	return nil
}

func printAccl(writer *tabwriter.Writer, item AcclItem) {
	fmt.Fprintf(writer, "%s\t%s\t%s\t%d\t%d\t%s\t%s\t%s\n",
		item.Name, item.Namespace, item.SchedulingGroupName, item.Priority, item.GPUs, item.ReconcileStatus, item.Status.Phase, item.Image)
}

type AcclHistory struct {
	Name string `json:"name"`
	// TODO: fix mam, need namespace
	// Namespace           string     `json:"namespace"`
	SchedulingGroupName string   `json:"schedulingGroup"`
	Priority            int      `json:"priority"`
	DeviceCount         int      `json:"deviceCount"`
	Status              string   `json:"status"`
	StartTime           string   `json:"startTime"`
	EndTime             string   `json:"endTime"`
	MafEnvs             []MafEnv `json:"mafEnvs"`
}

type MafEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func PrintAcclHistoryListK8sStyle(data io.ReadCloser) error {
	var acclHistoryItems []AcclHistory
	if err := json.NewDecoder(data).Decode(&acclHistoryItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(acclHistoryItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tSCHEDULING_GROUP_NAME\tPRIORITY\tGPUS\tSTATUS\tSTART_TIME\tEND_TIME")

	for _, item := range acclHistoryItems {
		printAcclHistory(writer, item)
	}

	return nil
}

func printAcclHistory(writer *tabwriter.Writer, item AcclHistory) {
	fmt.Fprintf(writer, "%s\t%s\t%d\t%d\t%s\t%s\t%s\n",
		item.Name, item.SchedulingGroupName, item.Priority, item.DeviceCount, item.Status, item.StartTime, item.EndTime)
}

func PrettyPrintJSON(rc io.ReadCloser) error {
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	var parsedData interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	} else if parsedData == nil {
		return fmt.Errorf("no resources found")
	}

	prettyData, err := json.MarshalIndent(parsedData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	fmt.Println(string(prettyData))
	return nil
}

type CpnItem struct {
	Name                string    `json:"name"`
	Namespace           string    `json:"namespace"`
	NodeName            string    `json:"nodeName"`
	SchedulingGroupName string    `json:"schedulingGroupName"`
	IP                  string    `json:"ip"`
	Status              CpnStatus `json:"status"`
	UpdatedAt           string    `json:"updatedAt"`
}

type CpnStatus struct {
	Phase          string `json:"phase"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

func PrintCpnK8sStyle(data io.ReadCloser) error {
	var acclItem *CpnItem
	if err := json.NewDecoder(data).Decode(&acclItem); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if acclItem == nil {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tNODE_NAME\tSCHEDULING_GROUP_NAME\tIP\tSTATUS\tUPDATED_AT")

	printCpn(writer, *acclItem)

	return nil
}

func PrintCpnListK8sStyle(data io.ReadCloser) error {
	var acclItems []CpnItem
	if err := json.NewDecoder(data).Decode(&acclItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(acclItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	if _, err := fmt.Fprintln(writer,
		"NAME\tNAMESPACE\tNODE_NAME\tSCHEDULING_GROUP_NAME\tIP\tSTATUS\tUPDATED_AT"); err != nil {
		return err
	}

	for _, item := range acclItems {
		printCpn(writer, item)
	}

	return nil
}

func printCpn(writer *tabwriter.Writer, item CpnItem) {
	if _, err := fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
		item.Name, item.Namespace, item.NodeName, item.SchedulingGroupName, item.IP, item.Status,
		item.UpdatedAt); err != nil {
		panic(err)
	}
}

type NodeFailureHistory struct {
	NodeName string `json:"nodeName"`
	// TODO: fix mam, need namespace
	// Namespace           string     `json:"namespace"`
	CheckType string `json:"checkType"`
	Time      string `json:"time"`
}

func PrintNodeFailureHistoryListK8sStyle(data io.ReadCloser) error {
	var nodeFailureHistoryItems []NodeFailureHistory
	if err := json.NewDecoder(data).Decode(&nodeFailureHistoryItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(nodeFailureHistoryItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NODE_NAME\tCHECK_TYPE\tTIME")

	for _, item := range nodeFailureHistoryItems {
		printNodeFailureHistory(writer, item)
	}

	return nil
}

func printNodeFailureHistory(writer *tabwriter.Writer, item NodeFailureHistory) {
	fmt.Fprintf(writer, "%s\t%s\t%s\n",
		item.NodeName, item.CheckType, item.Time)
}

type DeviceFailureHistory struct {
	NodeName string `json:"nodeName"`
	// TODO: fix mam, need namespace
	// Namespace           string     `json:"namespace"`
	DeviceIndex int    `json:"deviceIndex"`
	CheckType   string `json:"checkType"`
	Time        string `json:"time"`
}

func PrintDeviceFailureHistoryListK8sStyle(data io.ReadCloser) error {
	var deviceFailureHistoryItems []DeviceFailureHistory
	if err := json.NewDecoder(data).Decode(&deviceFailureHistoryItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(deviceFailureHistoryItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NODE_NAME\tDEVICE_INDEX\tCHECK_TYPE\tTIME")

	for _, item := range deviceFailureHistoryItems {
		printDeviceFailureHistory(writer, item)
	}

	return nil
}

func printDeviceFailureHistory(writer *tabwriter.Writer, item DeviceFailureHistory) {
	fmt.Fprintf(writer, "%s\t%d\t%s\t%s\n",
		item.NodeName, item.DeviceIndex, item.CheckType, item.Time)
}

type Flavor struct {
	Name                string   `json:"name"`
	SchedulingGroupName string   `json:"schedulingGroupName"`
	DeviceCount         int      `json:"deviceCount"`
	MafEnvs             []MafEnv `json:"mafEnvs,omitempty"`
}

func PrintFlavorK8sStyle(data io.ReadCloser) error {
	var flavorItem Flavor
	if err := json.NewDecoder(data).Decode(&flavorItem); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tSCHEDULING_GROUP_NAME\tGPUS")

	printFlavor(writer, flavorItem)

	return nil
}

func PrintFlavorListK8sStyle(data io.ReadCloser) error {
	var flavorItems []Flavor
	if err := json.NewDecoder(data).Decode(&flavorItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(flavorItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tSCHEDULING_GROUP_NAME\tGPUS")

	for _, item := range flavorItems {
		printFlavor(writer, item)
	}

	return nil
}

func printFlavor(writer *tabwriter.Writer, item Flavor) {
	fmt.Fprintf(writer, "%s\t%s\t%d\n",
		item.Name, item.SchedulingGroupName, item.DeviceCount)
}

type MafVersion struct {
	Tag     string   `json:"tag"`
	Image   string   `json:"image"`
	Enabled bool     `json:"enabled"`
	Latest  bool     `json:"latest"`
	MafEnvs []MafEnv `json:"mafEnvs"`
}

func PrintMafVersionK8sStyle(data io.ReadCloser) error {
	var mafVersionItem MafVersion
	if err := json.NewDecoder(data).Decode(&mafVersionItem); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "TAG\tIMAGE\tENABLED\tLATEST")

	printMafVersion(writer, mafVersionItem)

	return nil
}

func PrintMafVersionListK8sStyle(data io.ReadCloser) error {
	var mafVersionItems []MafVersion
	if err := json.NewDecoder(data).Decode(&mafVersionItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(mafVersionItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "TAG\tIMAGE\tENABLED\tLATEST")

	for _, item := range mafVersionItems {
		printMafVersion(writer, item)
	}

	return nil
}

func printMafVersion(writer *tabwriter.Writer, item MafVersion) {
	fmt.Fprintf(writer, "%s\t%s\t%t\t%t\n",
		item.Tag, item.Image, item.Enabled, item.Latest)
}

type SchedulingGroup struct {
	Name             string `json:"name"`
	Namespace        string `json:"namespace"`
	SchedulingPolicy string `json:"schedulingPolicy"`
	AllocationPolicy string `json:"allocationPolicy"`
}

func PrintSchedulingGroupK8sStyle(data io.ReadCloser) error {
	var schedulingGroup SchedulingGroup
	if err := json.NewDecoder(data).Decode(&schedulingGroup); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tSCHEDULING_POLICY\tALLOCATION_POLICY")

	printSchedulingGroup(writer, schedulingGroup)

	return nil
}

func PrintSchedulingGroupListK8sStyle(data io.ReadCloser) error {
	var schedulingGroupItems []SchedulingGroup
	if err := json.NewDecoder(data).Decode(&schedulingGroupItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(schedulingGroupItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tNAMESPACE\tSCHEDULING_POLICY\tALLOCATION_POLICY")

	for _, item := range schedulingGroupItems {
		printSchedulingGroup(writer, item)
	}

	return nil
}

func printSchedulingGroup(writer *tabwriter.Writer, item SchedulingGroup) {
	fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n",
		item.Name, item.Namespace, item.SchedulingPolicy, item.AllocationPolicy)
}

type SchedulerInfo struct {
	SchedulingGroupName string   `json:"schedulingGroupName"`
	AcceleratorOrder    []string `json:"acceleratorOrder"`
	UpdatedAt           string   `json:"updatedAt"`
}

func PrintSchedulerInfoK8sStyle(data io.ReadCloser) error {
	var schedulerInfoItem SchedulerInfo
	if err := json.NewDecoder(data).Decode(&schedulerInfoItem); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tACCELERATOR_ORDER\tUPDATED_AT")

	printSchedulerInfo(writer, schedulerInfoItem)

	return nil
}

func PrintSchedulerInfoListK8sStyle(data io.ReadCloser) error {
	var schedulerInfoItems []SchedulerInfo
	if err := json.NewDecoder(data).Decode(&schedulerInfoItems); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	if len(schedulerInfoItems) == 0 {
		return fmt.Errorf("no resources found")
	}

	writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NAME\tACCELERATOR_ORDER\tUPDATED_AT")

	for _, item := range schedulerInfoItems {
		printSchedulerInfo(writer, item)
	}

	return nil
}

func printSchedulerInfo(writer *tabwriter.Writer, item SchedulerInfo) {
	fmt.Fprintf(writer, "%s\t%v\t%s\n",
		item.SchedulingGroupName, item.AcceleratorOrder, item.UpdatedAt)
}

package get

import (
	"github.com/deviceinsight/kafkactl/operations"
	"github.com/deviceinsight/kafkactl/operations/consumergroups"
	"github.com/deviceinsight/kafkactl/operations/k8s"
	"github.com/deviceinsight/kafkactl/output"
	"github.com/spf13/cobra"
)

func newGetConsumerGroupsCmd() *cobra.Command {

	var flags consumergroups.GetConsumerGroupFlags

	var cmdGetConsumerGroups = &cobra.Command{
		Use:     "consumer-groups",
		Aliases: []string{"cg"},
		Short:   "list available consumerGroups",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if !(&k8s.K8sOperation{}).TryRun(cmd, args) {
				if err := (&consumergroups.ConsumerGroupOperation{}).GetConsumerGroups(flags); err != nil {
					output.Fail(err)
				}
			}
		},
	}

	cmdGetConsumerGroups.Flags().StringVarP(&flags.OutputFormat, "output", "o", flags.OutputFormat, "output format. One of: json|yaml|wide|compact")
	cmdGetConsumerGroups.Flags().StringVarP(&flags.FilterTopic, "topic", "t", "", "show groups for given topic only")

	if err := cmdGetConsumerGroups.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations.CompleteTopicNames(cmd, args, toComplete)
	}); err != nil {
		panic(err)
	}

	return cmdGetConsumerGroups
}

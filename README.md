# poc-neon
Project Neon

# Run tests Command
rover test \
      -b <path to tests folder> \
      -env <environment name> \
      -level <level> \
      -tfstate <name of deployed state file> \
      -d

## Example Run Tests Command
./rover/rover.sh ./caf/test \
      -b ./caf/test/azure \
      -env crimes3_demo \
      -level level0 \
      -tfstate launchpad.tfstate \
      -d
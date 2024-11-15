<?php
declare(strict_types=1);

namespace App\Command\Day1;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:1:1', description: 'Challenge Day 1 - Part 1')]
class PartOneCommand extends Command
{
    public function __construct(
        private readonly FileReaderService $fileReaderService,
        protected readonly ?string $name = null,
    ) {
        parent::__construct($name);
    }

    protected function configure(): void
    {
        $this
            ->addOption(
                'input',
                'i',
                InputOption::VALUE_REQUIRED,
                'File name to use as input',
                'example'
            );
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        /** @var string $inputFile */
        $inputFile = $input->getOption('input');

        $lines = $this->fileReaderService->readLines(1, $inputFile . '.txt');

        $mostCalories = 0;

        while (count($lines) > 0) {
            $calories = self::getNextElfCalories($lines);

            if ($calories > $mostCalories) {
                $mostCalories = $calories;
            }
        }

        $output->writeln('Most Calories: ' . $mostCalories);

        return Command::SUCCESS;
    }

    /**
     * @param string[] $lines
     */
    private static function getNextElfCalories(array &$lines): int
    {
        $next = array_shift($lines);

        if (!$next) {
            return 0;
        }

        return (int) $next + self::getNextElfCalories($lines);
    }
}

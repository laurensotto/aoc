<?php
declare(strict_types=1);

namespace App\Command\Day4;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:4:2', description: 'Challenge Day 4 - Part 2')]
class PartTwoCommand extends Command
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

        $lines = $this->fileReaderService->readLines(4, $inputFile . '.txt');

        $fullOverlapCount = 0;

        foreach ($lines as $line) {
            $values = explode(',', $line);

            $firstElfSections  = self::getElfSections($values[0]);
            $secondElfSections = self::getElfSections($values[1]);

            $intersection = array_intersect($firstElfSections, $secondElfSections);

            if ($intersection !== []) {
                $fullOverlapCount++;
            }
        }

        $output->writeln('Result: ' . $fullOverlapCount);

        return Command::SUCCESS;
    }

    /** @return int[] */
    private static function getElfSections(string $values): array
    {
        $splitValues = explode('-', $values);

        return range((int) $splitValues[0], (int) $splitValues[1]);
    }
}

<?php
declare(strict_types=1);

namespace App\Command\Day3;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:3:1', description: 'Challenge Day 3 - Part 1')]
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

        $lines = $this->fileReaderService->readLines(3, $inputFile . '.txt');

        $letters      = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
        $priorityList = [];
        foreach (str_split($letters) as $index => $letter) {
            $priorityList[$letter] = $index + 1;
        }


        $priorityCount = 0;
        foreach ($lines as $line) {
            $items       = str_split($line);
            $amountItems = count($items);

            if ($amountItems <= 0) {
                continue;
            }

            /** @var int<1, max> $compartmentSize */
            $compartmentSize = $amountItems / 2;

            $compartments = array_chunk($items, $compartmentSize);

            $overlap = array_intersect($compartments[0], $compartments[1]);

            if ($value = array_pop($overlap)) {
                $priorityCount += $priorityList[$value];
            }
        }

        $output->writeln('Result: ' . $priorityCount);

        return Command::SUCCESS;
    }
}

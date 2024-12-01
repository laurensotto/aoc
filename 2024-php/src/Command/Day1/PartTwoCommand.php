<?php
declare(strict_types=1);

namespace App\Command\Day1;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:1:2', description: 'Challenge Day 1 - Part 2')]
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

        $lines = $this->fileReaderService->readLines(1, $inputFile . '.txt');

        /** @var int[] $listOne */
        $listOne = [];
        /** @var int[] $listTwo */
        $listTwo = [];

        foreach ($lines as $line) {
            $values = explode('   ', $line);

            $listOne[] = (int) $values[0];
            $listTwo[] = (int) $values[1];
        }

        $total = 0;

        foreach ($listOne as $value) {
            $total += self::getSimilarityScore($value, $listTwo);
        }

        $output->writeln('Result: ' . $total);

        return Command::SUCCESS;
    }

    /** @param int[] $listTwo */
    private static function getSimilarityScore(int $valueFromListOne, array $listTwo): int
    {
        $values = array_filter($listTwo, fn (int $value) => $value === $valueFromListOne);

        return $valueFromListOne * count($values);
    }
}

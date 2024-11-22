<?php
declare(strict_types=1);

namespace App\Command\Day5;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:5:1', description: 'Challenge Day 5 - Part 1')]
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

        $lines = $this->fileReaderService->readLines(5, $inputFile . '.txt');

        $endIndexInitialStacks = self::getEndIndexInitialStacks($lines);

        $stacks = self::getStartingStacks($lines, $endIndexInitialStacks);

        for ($i = $endIndexInitialStacks + 1; $i <= count($lines) - 1; $i++) {
            $move = self::getMove($lines[$i]);

            $amount = $move[0];
            $from   = $move[1] - 1;
            $to     = $move[2] - 1;

            for ($j = 0; $j < $amount; $j++) {
                $stacks[$to][] = array_pop($stacks[$from]);
            }
        }

        $answer = '';
        foreach ($stacks as $stack) {
            $answer .= array_pop($stack);
        }

        $output->writeln('Result: ' . $answer);

        return Command::SUCCESS;
    }


    /**
     * @param string[] $lines
     * @return array<int, string[]>
     */
    private static function getStartingStacks(array $lines, int $endIndexInitialStacks): array
    {
        $width = strlen($lines[0]);

        $columns = (($width + 1) / 4);

        $startingStacks = [];

        for ($i = 0; $i < $columns; $i++) {
            $startingStacks[] = [];
        }

        for ($i = $endIndexInitialStacks - 2; $i >= 0; $i--) {
            for ($j = 1; $j < $width; $j += 4) {
                $stackIndex = ($j - 1) / 4;

                if ($lines[$i][$j] !== ' ') {
                    $startingStacks[$stackIndex][] = $lines[$i][$j];
                }
            }
        }

        return $startingStacks;
    }

    /** @param string[] $lines */
    private static function getEndIndexInitialStacks(array $lines): int
    {
        $endIndex = 0;

        foreach ($lines as $index => $line) {
            if ($line === '') {
                $endIndex = $index;
            }
        }

        return $endIndex;
    }

    /** @return int[] */
    private static function getMove(string $line): array
    {
        preg_match_all('/\d+/', $line, $matches);

        return array_map('intval', $matches[0]);
    }
}
